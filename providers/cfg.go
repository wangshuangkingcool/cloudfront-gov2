package providers

import (
	"log"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func Hello() string {
	return "Config"
}

func GetConfigForProfile(profile string) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), 
    	config.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	return cfg
}

func GetConfigFromCreds() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), 
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("", "", "")))

	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	return cfg

}