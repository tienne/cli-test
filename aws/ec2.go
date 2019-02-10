package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

func EC2List() {
	sessionConfig := SessionConfig{}

	awsSession, _ := NewSession(sessionConfig)
	service := ec2.New(awsSession)

	instances, _ := service.DescribeInstances(&ec2.DescribeInstancesInput{})

	for index, instance := range instances.Reservations {
		fmt.Println(instance.Instances)
		fmt.Println(index)
		os.Exit(1)
	}

}
