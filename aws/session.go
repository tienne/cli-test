package aws

import (
	"fmt"
	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type SessionConfig struct {
	AccessKey string
	SecretKey string
	Region    string
}

func NewSession(config SessionConfig) (*session.Session, error) {
	awsConfig := sdk.NewConfig()

	if config.Region != "" {
		awsConfig.Region = sdk.String(config.Region)
	}
	fmt.Println(config.AccessKey)
	fmt.Println(config.SecretKey)
	if config.AccessKey != "" || config.SecretKey != "" {
		// awsConfig.Credentials = credentials.NewStaticCredentials()
	}

	awsSession, err := session.NewSession(awsConfig)
	// Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),

	if err != nil {
		return nil, err
	}

	return awsSession, nil
}
