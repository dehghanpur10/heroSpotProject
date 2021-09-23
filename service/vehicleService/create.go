package vehicleService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)


func (s *vehicleService) CreateVehicle(vehicle models.Vehicle) error {
	item, err := dynamodbattribute.MarshalMap(vehicle)
	if err != nil {
		fmt.Print("vehicleService.CreateVehicle - MarshalMap - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("VehicleSpot"),
	}
	_, err = s.db.PutItem(input)
	if err != nil {
		fmt.Print("vehicleService.CreateVehicle - PutItem - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	return nil
}

