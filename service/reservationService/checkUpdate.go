package reservationService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *ReservationService) CheckUpdate(reservationId string) (models.Reservation, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String("ReservationSpot"),
		Key: map[string]*dynamodb.AttributeValue{
			"reservation_id": &dynamodb.AttributeValue{
				S: aws.String(reservationId),
			},
		},
	}

	result, err := s.db.GetItem(getItemInput)
	if err != nil {
		err = fmt.Errorf("reservationService.cehckUpdate - %v - %v", lib.ErrInternal, err)
		return models.Reservation{}, err
	}

	if result.Item == nil {
		err = fmt.Errorf("reservationService.cehckUpdate - %v", lib.ErrNotFound)
		return models.Reservation{}, err
	}

	var reservation models.Reservation
	err = dynamodbattribute.UnmarshalMap(result.Item, &reservation)
	if err != nil {
		err = fmt.Errorf("reservationService.cehckUpdate - unmarshalMap - %v - %v", lib.ErrInternal, err)
		return models.Reservation{}, err
	}

	return reservation, nil

}
