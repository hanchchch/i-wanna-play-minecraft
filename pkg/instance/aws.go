package instance

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type AWSInstance struct {
	InstanceID string
	EC2        *ec2.Client
}

type AWSInstanceOptions struct {
	InstanceID string
	Region     string
	Profile    string
}

func NewAWSInstance(o AWSInstanceOptions) *AWSInstance {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(o.Profile))
	cfg.Region = o.Region
	if err != nil {
		panic(err)
	}
	return &AWSInstance{
		InstanceID: o.InstanceID,
		EC2:        ec2.NewFromConfig(cfg),
	}
}

func (i *AWSInstance) Start() (*ec2.StartInstancesOutput, error) {
	return i.EC2.StartInstances(context.TODO(), &ec2.StartInstancesInput{
		InstanceIds: []string{i.InstanceID},
	})
}

func (i *AWSInstance) Stop() (*ec2.StopInstancesOutput, error) {
	return i.EC2.StopInstances(context.TODO(), &ec2.StopInstancesInput{
		InstanceIds: []string{i.InstanceID},
	})
}

func (i *AWSInstance) RunFor(d time.Duration) error {
	if _, err := i.Start(); err != nil {
		return err
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-time.After(d):
		fmt.Println("shutting down, stopping instance")
		if _, err := i.Stop(); err != nil {
			return err
		}
	case <-c:
		fmt.Println("stopping instance")
		if _, err := i.Stop(); err != nil {
			return err
		}
	}

	return nil
}
