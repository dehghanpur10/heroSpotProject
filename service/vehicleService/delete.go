package vehicleService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"spotHeroProject/lib"
)

func (s *vehicleService) DeleteVehicle(id string) error {
	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(lib.VEHICLE_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"vehicle_id": {
				S: aws.String(id),
			},
		},
	}
	_, err := s.db.DeleteItem(deleteItemInput)
	if err != nil {
		fmt.Print("vehicleService.deleteVehicle - DeleteItem - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	return nil
}

