package config

import (
	"github.com/mozzet/cli/env"
)

type Config struct {
	env.Env
	AwsRegion    string
	AwsAccessKey string
	AwsSecretKey string
}
