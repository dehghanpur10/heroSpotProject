package lib

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDynamoDB() (*dynamodb.DynamoDB, error) {
	region := os.Getenv("AWS_REGION")
	accessToken := os.Getenv("ACCESS_TOKEN")
	secretKey := os.Getenv("SECRET_KEY")
	if region == "" || accessToken == "" || secretKey == "" {
		return &dynamodb.DynamoDB{}, fmt.Errorf("GetdynamoDb - env - %v", ErrNotFound)
	}
	credential := credentials.NewStaticCredentials(accessToken, secretKey, "")
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credential,
	},
	)
	if err != nil {
		return &dynamodb.DynamoDB{}, fmt.Errorf("GetdynamoDb - newSession - %v", err)
	}
	return dynamodb.New(awsSession), nil
}
