package reservationService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"spotHeroProject/lib"
)

func (s *ReservationService) DeleteReservationService(reservationId string) error {
	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(lib.RESERVATION_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"reservation_id": {
				S: aws.String(reservationId),
			},
		},
	}
	_, err := s.db.DeleteItem(deleteItemInput)
	if err != nil {
		fmt.Print("reservationService.delete - deleteItem - ")
		return fmt.Errorf("%w - %v", lib.ErrInternal, err)
	}
	return nil
}
