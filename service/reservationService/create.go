package reservationService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *ReservationService) CreateReservation(reservation models.Reservation) error {
	item, err := dynamodbattribute.MarshalMap(reservation)
	if err != nil {
		fmt.Print("reservationService.Create - marshalMap - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(lib.RESERVATION_TABLE_NAME),
	}
	_, err = s.db.PutItem(input)
	if err != nil {
		fmt.Printf("reservationService.Create - putItem - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	return nil
}
