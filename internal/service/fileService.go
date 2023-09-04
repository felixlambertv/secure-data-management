package service

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/s3"
	"secure-data-management/internal/controller/requests"
	"secure-data-management/internal/model"
	"secure-data-management/internal/repository"
	"strings"
	"time"
)

type FileService interface {
	UploadFile(request requests.UploadFileRequest, userId string, idToken string) error
}

type S3FileService struct {
	userRepository repository.UserRepository
	fileRepository repository.FileRepository
	sess           *session.Session
}

func NewS3FileService(userRepository repository.UserRepository, fileRepository repository.FileRepository, sess *session.Session) *S3FileService {
	return &S3FileService{userRepository: userRepository, fileRepository: fileRepository, sess: sess}
}

func (s *S3FileService) UploadFile(request requests.UploadFileRequest, userId string, idToken string) error {
	user, err := s.userRepository.FindById(context.Background(), userId)
	ci := cognitoidentity.New(s.sess)
	credential, err := ci.GetCredentialsForIdentity(&cognitoidentity.GetCredentialsForIdentityInput{
		IdentityId: aws.String(user.PoolId),
		Logins: map[string]*string{
			"cognito-idp.ap-southeast-1.amazonaws.com/ap-southeast-1_zFFMzwFGK": aws.String(idToken),
		},
	})
	if err != nil {
		return err
	}

	userSess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(
			*credential.Credentials.AccessKeyId,
			*credential.Credentials.SecretKey,
			*credential.Credentials.SessionToken,
		),
	})
	if err != nil {
		return err
	}

	s3Svc := s3.New(userSess)
	for _, f := range request.Files {
		file, err := f.Open()
		originalFileName := f.Filename
		filenameWithOutSpace := strings.Replace(originalFileName, " ", "", len(originalFileName))
		fileName := time.Now().Format("200601020405") + "-" + filenameWithOutSpace
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = s3Svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String("go-file-upload"),
			Key:    aws.String("uploads/" + fileName),
			Body:   file,
		})
		if err != nil {
			return err
		}

		metadata := model.Metadata{
			Name:     originalFileName,
			Size:     int(request.Files[0].Size),
			UploadAt: int(time.Now().Unix()),
			Type:     getFileExtension(f.Filename),
		}
		fi := &model.File{
			Permission: request.Permission,
			Metadata:   metadata,
		}

		err = s.fileRepository.Save(context.Background(), fi)
		if err != nil {
			return err
		}
	}

	return nil
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
