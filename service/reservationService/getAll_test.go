package reservationService

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"spotHeroProject/mocks"
	"spotHeroProject/models"
	"testing"
)

func TestGetAllReservation(t *testing.T) {
	items := []map[string]*dynamodb.AttributeValue{
		{
			"reservation_id":&dynamodb.AttributeValue{
				S: aws.String("2"),
			},
			"update_possible":&dynamodb.AttributeValue{
				BOOL: aws.Bool(false),
			},
		},
	}
	tests := []struct {
		name                string
		scanError           error
		scanResult          *dynamodb.ScanOutput
		expectedReservation []models.Reservation
	}{
		{name:"scan error",scanError: errors.New("scan thrown an error")},
		{name:"ok",scanResult: &dynamodb.ScanOutput{Items: items},expectedReservation: []models.Reservation{{Id: "2",UpdatePossible: false}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := new(mocks.DynamoDBAPI)
			service := New(db)
			db.On("Scan", mock.Anything).Return(test.scanResult, test.scanError)

			reservation, err := service.GetAll()

			if err != nil {
				assert.Contains(t, err.Error(), test.scanError.Error())
			} else {
				assert.Nil(t, test.scanError)
			}

			assert.Equal(t, test.expectedReservation, reservation)
		})
	}
}
