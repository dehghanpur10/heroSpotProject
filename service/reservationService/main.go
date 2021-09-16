package reservationService

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

type ReservationService struct {
	db dynamodbiface.DynamoDBAPI
}

func New(db dynamodbiface.DynamoDBAPI) *ReservationService {
	return &ReservationService{
		db: db,
	}
}

