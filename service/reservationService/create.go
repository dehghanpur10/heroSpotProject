package reservationService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *ReservationService) Create(reservation models.Reservation) error {
	item, err := dynamodbattribute.MarshalMap(reservation)
	if err != nil {
		return fmt.Errorf("reservationService.Create - marshaoMap - %w - %v", lib.ErrInternal, err)
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("ReservationSpot"),
	}
	_, err = s.db.PutItem(input)
	if err != nil {
		return fmt.Errorf("reservationService.Create - putItem - %w - %v", lib.ErrInternal, err)
	}
	return nil
}
