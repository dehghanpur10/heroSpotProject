package reservationService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *ReservationService) GetAllReservation() ([]models.Reservation, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(lib.RESERVATION_TABLE_NAME),
	}
	result, err := s.db.Scan(scanInput)
	if err != nil {
		fmt.Print("reservationService.GetAll - Scan - ")
		return nil, fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}

	var reservations []models.Reservation
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &reservations)
	if err != nil {
		fmt.Print("reservationService.GetAll  - UnmarshalListOfMaps - ")
		return nil, fmt.Errorf("%w - %v", lib.ErrInternal, err)

	}

	return reservations, nil
}
