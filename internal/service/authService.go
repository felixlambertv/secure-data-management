package service

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"secure-data-management/config"
	"secure-data-management/internal/controller/requests"
	"secure-data-management/internal/model"
	"secure-data-management/internal/repository"
)

type AuthService interface {
	Register(request requests.UserRegisterRequest) error
	Login(request requests.LoginRequest) (*cognitoidentityprovider.InitiateAuthOutput, error)
	Verify(request requests.AccountVerifyRequest) error
}

type CognitoAuthService struct {
	userRepo       repository.UserRepository
	userPoolId     string
	appClientId    string
	cognitoService *cognitoidentityprovider.CognitoIdentityProvider
	sess           *session.Session
}

func NewCognitoAuthService(userRepo repository.UserRepository, sess *session.Session, config *config.AWSConfig) *CognitoAuthService {
	cognitoService := cognitoidentityprovider.New(sess)
	return &CognitoAuthService{
		userRepo:       userRepo,
		userPoolId:     config.CognitoUserPoolId,
		appClientId:    config.CognitoClientId,
		cognitoService: cognitoService,
		sess:           sess,
	}
}

func (c *CognitoAuthService) Register(request requests.UserRegisterRequest) error {
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(c.appClientId),
		Password: aws.String(request.Password),
		Username: aws.String(request.Username),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(request.Email),
			},
		},
	}

	output, err := c.cognitoService.SignUp(signUpInput)
	if err != nil {
		return err
	}

	user := &model.User{
		ID:         *output.UserSub,
		Username:   request.Username,
		Email:      request.Email,
		IsVerified: *output.UserConfirmed,
	}
	err = c.userRepo.Save(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (c *CognitoAuthService) Login(request requests.LoginRequest) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeUserPasswordAuth),
		ClientId: aws.String(c.appClientId),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(request.Username),
			"PASSWORD": aws.String(request.Password),
		},
	}

	auth, err := c.cognitoService.InitiateAuth(authInput)
	if err != nil {
		return nil, err
	}

	user, err := c.userRepo.FindByUsername(context.Background(), request.Username)
	if err != nil {
		return nil, err
	}

	if len(user.PoolId) == 0 {
		ci := cognitoidentity.New(c.sess)
		id, err := ci.GetId(&cognitoidentity.GetIdInput{
			IdentityPoolId: aws.String(user.PoolId),
			Logins: map[string]*string{
				"cognito-idp.ap-southeast-1.amazonaws.com/" + c.userPoolId: auth.AuthenticationResult.IdToken,
			},
		})
		if err != nil {
			return nil, err
		}
		user.PoolId = *id.IdentityId
		err = c.userRepo.Update(context.Background(), user)
	}

	if err != nil {
		return nil, err
	}
	return auth, err
}

func (c *CognitoAuthService) Verify(request requests.AccountVerifyRequest) error {
	user, err := c.userRepo.FindByUsername(context.Background(), request.Username)
	if err != nil {
		return err
	}

	signUpInput := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(c.appClientId),
		ConfirmationCode: aws.String(request.Code),
		Username:         aws.String(request.Username),
	}

	_, err = c.cognitoService.ConfirmSignUp(signUpInput)
	if err != nil {
		return err
	}

	user.IsVerified = true
	err = c.userRepo.Update(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
