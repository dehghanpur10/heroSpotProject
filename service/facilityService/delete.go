package facilityService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"spotHeroProject/lib"
)

func (s *FacilityService) DeleteFacilityService(facilityId string) error {
	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(lib.FACILITY_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"facility_id": {
				S: aws.String(facilityId),
			},
		},
	}
	_, err := s.db.DeleteItem(deleteItemInput)
	if err != nil {
		fmt.Print("facilityService.delete - deleteItem - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	return nil
}
