package reservationService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *ReservationService) GetAll() ([]models.Reservation, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String("ReservationSpot"),
	}
	result, err := s.db.Scan(scanInput)
	if err != nil {
		return nil, fmt.Errorf("reservationService.GetAll - Scan - %w - %v", lib.ErrInternal, err)
	}

	var reservations []models.Reservation
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &reservations)
	if err != nil {
		err = fmt.Errorf("reservationService.GetAll  - UnmarshalListOfMaps - %w - %v", lib.ErrInternal, err)
		return nil, err
	}

	return reservations, nil
}
