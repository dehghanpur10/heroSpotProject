package reservationService

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"spotHeroProject/mocks"
	"testing"
)

func TestDeleteReservationService(t *testing.T) {
	tests := []struct {
		name            string
		deleteItemError error
	}{
		{name: "ok"},
		{name: "deleteItem error", deleteItemError: errors.New("deleteItem error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//Arrange
			mockDB := new(mocks.DynamoDBAPI)
			service := New(mockDB)
			mockDB.On("DeleteItem", mock.Anything).Return(&dynamodb.DeleteItemOutput{}, test.deleteItemError)
			//Act
			err := service.DeleteReservationService("")
			//Assert
			if err != nil {
				assert.Contains(t, err.Error(), test.deleteItemError.Error())
			} else {
				assert.Nil(t, test.deleteItemError)
			}
		})
	}
}
