package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWSConfig struct {
	Region            string
	AccessKey         string
	SecretKey         string
	CognitoUserPoolId string
	CognitoClientId   string
}

func NewAWSConfig(region string, accessKey string, secretKey string, cognitoUserPoolId string, cognitoClientId string) *AWSConfig {
	return &AWSConfig{
		Region:            region,
		AccessKey:         accessKey,
		SecretKey:         secretKey,
		CognitoUserPoolId: cognitoUserPoolId,
		CognitoClientId:   cognitoClientId,
	}
}

func (c AWSConfig) NewSession() (*session.Session, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(c.Region),
		Credentials: credentials.NewStaticCredentials(c.AccessKey, c.SecretKey, ""),
	}
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
