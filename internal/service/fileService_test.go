package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"reflect"
	"secure-data-management/internal/controller/requests"
	"secure-data-management/internal/repository"
	"testing"
)

func TestNewS3FileService(t *testing.T) {
	type args struct {
		userRepository repository.UserRepository
		fileRepository repository.FileRepository
		sess           *session.Session
	}
	tests := []struct {
		name string
		args args
		want *S3FileService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewS3FileService(tt.args.userRepository, tt.args.fileRepository, tt.args.sess); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewS3FileService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS3FileService_GetMedia(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
		fileRepository repository.FileRepository
		sess           *session.Session
	}
	type args struct {
		fileId string
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &S3FileService{
				userRepository: tt.fields.userRepository,
				fileRepository: tt.fields.fileRepository,
				sess:           tt.fields.sess,
			}
			got, err := s.GetMedia(tt.args.fileId, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMedia() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS3FileService_UploadFile(t *testing.T) {
	type fields struct {
		userRepository repository.UserRepository
		fileRepository repository.FileRepository
		sess           *session.Session
	}
	type args struct {
		request requests.UploadFileRequest
		userId  string
		idToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &S3FileService{
				userRepository: tt.fields.userRepository,
				fileRepository: tt.fields.fileRepository,
				sess:           tt.fields.sess,
			}
			if err := s.UploadFile(tt.args.request, tt.args.userId, tt.args.idToken); (err != nil) != tt.wantErr {
				t.Errorf("UploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		arr    []string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFileExtension(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileExtension(tt.args.filename); got != tt.want {
				t.Errorf("getFileExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
