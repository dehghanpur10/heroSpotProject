package facilityService

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

type FacilityService struct {
	db dynamodbiface.DynamoDBAPI
}

func New(db dynamodbiface.DynamoDBAPI) *FacilityService {
	return &FacilityService{
		db: db,
	}
}

