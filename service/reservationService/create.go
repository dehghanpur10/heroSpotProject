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
		return fmt.Errorf("reservationService.Create - marshaoMap - %v - %v", lib.ErrInternal, err)
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Reservation"),
	}
	_, err = s.db.PutItem(input)
	if err != nil {
		return fmt.Errorf("reservationService.Create - putItem - %v - %v", lib.ErrInternal, err)
	}
	return nil
}
