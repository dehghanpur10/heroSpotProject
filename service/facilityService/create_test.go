package facilityService

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"spotHeroProject/mocks"
	"spotHeroProject/models"
	"testing"
)

func TestCreateFacilityService(t *testing.T) {
	tests := []struct{
		name string
		putItemError error
	}{
		{name: "ok"},
		{name: "putItem error",putItemError: errors.New("putItem error")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			mockDB := new(mocks.DynamoDBAPI)
			service := New(mockDB)
			mockDB.On("PutItem",mock.Anything).Return(&dynamodb.PutItemOutput{},test.putItemError)
			// Act
			err := service.CreateFacilityService(models.Facility{})
			// Assert
			if err != nil {
				assert.Contains(t, err.Error(),test.putItemError.Error())
			} else {
				assert.Nil(t, test.putItemError)
			}
		})
	}
}