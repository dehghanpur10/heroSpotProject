package vehicleService

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

type vehicleService struct {
	db dynamodbiface.DynamoDBAPI
}
func New(db dynamodbiface.DynamoDBAPI) *vehicleService {
	return &vehicleService{
		db: db,
	}
}