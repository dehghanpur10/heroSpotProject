package reservationService

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"spotHeroProject/mocks"
	"spotHeroProject/models"
	"testing"
)

func TestCreateReservation(t *testing.T) {
	tests := []struct {
		name         string
		putItemError error
	}{
		{name: "ok"},
		{name: "putItem error", putItemError: errors.New("putItem Error")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			db := new(mocks.DynamoDBAPI)
			service := New(db)
			db.On("PutItem", mock.Anything).Return(&dynamodb.PutItemOutput{}, test.putItemError)
			// Act
			err := service.CreateReservation(models.Reservation{})
			// Assert
			if err != nil {
				assert.Contains(t, err.Error(), test.putItemError.Error())
			} else {
				assert.Nil(t, test.putItemError)
			}
		})
	}
}
