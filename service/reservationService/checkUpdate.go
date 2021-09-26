package reservationService

import (
	"fmt"
	"spotHeroProject/lib"
	"spotHeroProject/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (s *ReservationService) CheckUpdateReservation(reservationId string) (models.Reservation, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(lib.RESERVATION_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"reservation_id": {
				S: aws.String(reservationId),
			},
		},
	}

	result, err := s.db.GetItem(getItemInput)
	if err != nil {
		fmt.Print("reservationService.checkUpdate - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return models.Reservation{}, err
	}

	if result.Item == nil {
		fmt.Print("reservationService.checkUpdate - ")
		err = fmt.Errorf("%w", lib.ErrNotFound)
		return models.Reservation{}, err
	}

	var reservation models.Reservation
	err = dynamodbattribute.UnmarshalMap(result.Item, &reservation)
	if err != nil {

		fmt.Print("reservationService.checkUpdate - unmarshalMap - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return models.Reservation{}, err
	}

	return reservation, nil
}
