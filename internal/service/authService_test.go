package service

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"reflect"
	"secure-data-management/config"
	"secure-data-management/internal/controller/requests"
	"secure-data-management/internal/repository"
	"testing"
)

func TestCognitoAuthService_Login(t *testing.T) {
	type fields struct {
		userRepo       repository.UserRepository
		userPoolId     string
		appClientId    string
		cognitoService *cognitoidentityprovider.CognitoIdentityProvider
		sess           *session.Session
	}
	type args struct {
		request requests.LoginRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *cognitoidentityprovider.InitiateAuthOutput
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CognitoAuthService{
				userRepo:       tt.fields.userRepo,
				userPoolId:     tt.fields.userPoolId,
				appClientId:    tt.fields.appClientId,
				cognitoService: tt.fields.cognitoService,
				sess:           tt.fields.sess,
			}
			got, err := c.Login(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCognitoAuthService_Register(t *testing.T) {
	type fields struct {
		userRepo       repository.UserRepository
		userPoolId     string
		appClientId    string
		cognitoService *cognitoidentityprovider.CognitoIdentityProvider
		sess           *session.Session
	}
	type args struct {
		request requests.UserRegisterRequest
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
			c := &CognitoAuthService{
				userRepo:       tt.fields.userRepo,
				userPoolId:     tt.fields.userPoolId,
				appClientId:    tt.fields.appClientId,
				cognitoService: tt.fields.cognitoService,
				sess:           tt.fields.sess,
			}
			if err := c.Register(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCognitoAuthService_Verify(t *testing.T) {
	type fields struct {
		userRepo       repository.UserRepository
		userPoolId     string
		appClientId    string
		cognitoService *cognitoidentityprovider.CognitoIdentityProvider
		sess           *session.Session
	}
	type args struct {
		request requests.AccountVerifyRequest
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
			c := &CognitoAuthService{
				userRepo:       tt.fields.userRepo,
				userPoolId:     tt.fields.userPoolId,
				appClientId:    tt.fields.appClientId,
				cognitoService: tt.fields.cognitoService,
				sess:           tt.fields.sess,
			}
			if err := c.Verify(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewCognitoAuthService(t *testing.T) {
	type args struct {
		userRepo repository.UserRepository
		sess     *session.Session
		config   *config.AWSConfig
	}
	tests := []struct {
		name string
		args args
		want *CognitoAuthService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCognitoAuthService(tt.args.userRepo, tt.args.sess, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCognitoAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}
