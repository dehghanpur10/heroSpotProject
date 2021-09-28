package facilityService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *FacilityService) CreateFacilityService(facility models.Facility) error{
	item, err := dynamodbattribute.MarshalMap(facility)
	if err != nil {
		fmt.Print("facilityService.Create - marshalMap - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(lib.FACILITY_TABLE_NAME),
	}
	_, err = s.db.PutItem(input)
	if err != nil {
		fmt.Printf("facilityService.Create - putItem - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	return nil
}
