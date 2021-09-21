package reservationService

import (
	"fmt"
	"spotHeroProject/lib"
	"spotHeroProject/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

/*
 TODO (@ryan.bm): The name of your function should be more descriptive,
 for example here I don't know if CheckUpdate return Facility updates or Reservation or ...

 use CheckReservationUpdate(), (remember to update the name in logs as well)
*/
func (s *ReservationService) CheckUpdate(reservationId string) (models.Reservation, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String("ReservationSpot"),
		Key: map[string]*dynamodb.AttributeValue{
			// TODO (@ryan.bm): fix lint errors (remove extra &dynamodb.AttributeValue)
			"reservation_id": &dynamodb.AttributeValue{
				S: aws.String(reservationId),
			},
		},
	}

	result, err := s.db.GetItem(getItemInput)
	if err != nil {
		// TODO (@ryan.bm): fix type (cehck) in logs
		err = fmt.Errorf("reservationService.cehckUpdate - %w - %v", lib.ErrInternal, err)
		return models.Reservation{}, err
	}

	if result.Item == nil {
		err = fmt.Errorf("reservationService.cehckUpdate - %w", lib.ErrNotFound)
		return models.Reservation{}, err
	}

	var reservation models.Reservation
	err = dynamodbattribute.UnmarshalMap(result.Item, &reservation)
	if err != nil {
		/*
			TODO (@ryan.bm)
			Why lib.ErrInternal ?
			Its better to just return the err without any changes, and fmt.Print() the error stack
			like this:

			fmt.Print("ReservationService.CheckReservationUpdate-UnmarshalMap-")
			return models.Reservation{}, err
		*/
		err = fmt.Errorf("reservationService.cehckUpdate - unmarshalMap - %w - %v", lib.ErrInternal, err)
		return models.Reservation{}, err
	}

	return reservation, nil
}
