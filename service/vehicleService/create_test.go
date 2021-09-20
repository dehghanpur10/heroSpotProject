package vehicleService

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"spotHeroProject/mocks"
	"spotHeroProject/models"
	"testing"
)

func TestCreateVehicleService(t *testing.T) {
	tests := []struct {
		name          string
		putItemError  error
		expectedError error
	}{
		{name: "ok"},
		{name: "putItemError", expectedError: errors.New("PutItem error"), putItemError: errors.New("PutItem error")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := new(mocks.DynamoDBAPI)
			db.On("PutItem", mock.Anything).Return(&dynamodb.PutItemOutput{}, test.putItemError)
			service := New(db)

			err := service.Create(models.Vehicle{})
			if err != nil {
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.Nil(t, test.expectedError)
			}

		})
	}

}
