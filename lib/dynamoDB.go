package lib

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetDynamoDB() (*dynamodb.DynamoDB, error) {

	if AWS_REGION == "" || ACCESS_TOKEN == "" || SECRET_KEY == "" {
		return &dynamodb.DynamoDB{}, fmt.Errorf("GetdynamoDb - env - %v", ErrNotFound)
	}
	credential := credentials.NewStaticCredentials(ACCESS_TOKEN, SECRET_KEY, "")
	awsSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Credentials: credential,
	},
	)
	if err != nil {
		return &dynamodb.DynamoDB{}, fmt.Errorf("GetdynamoDb - newSession - %v", err)
	}
	return dynamodb.New(awsSession), nil
}
