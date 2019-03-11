package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

type Region string

func EC2List() {
	sessionConfig := SessionConfig{}

	if sessionConfig.Region == "" {
		sessionConfig.Region = "ap-northeast-2"
	}

	awsSession, _ := NewSession(sessionConfig)
	service := ec2.New(awsSession)

	limit := int64(1000)

	instances, _ := service.DescribeInstances(&ec2.DescribeInstancesInput{
		MaxResults: &limit,
	})

	for index, instance := range instances.Reservations {
		fmt.Println(instance.Instances)
		fmt.Println(index)
		os.Exit(1)
	}

}
