package searchService

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"spotHeroProject/mocks"
	"spotHeroProject/models"
	"testing"
)


func TestGetWithLatAndLon(t *testing.T) {
	tests := []struct {
		name             string
		scanError        error
		queryOutput       *dynamodb.QueryOutput
		expectedFacility []models.Facility
		expectedError    error
	}{
		{name: "ok", queryOutput: &dynamodb.QueryOutput{Items: item}, expectedFacility: facilities},
		{name: "scan error", scanError: errors.New("scan error"), expectedError: errors.New("scan error")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db := new(mocks.DynamoDBAPI)
			service := NewSearchService(db)
			db.On("Query", mock.Anything).Return(test.queryOutput, test.scanError)

			facility, err := service.GetFacilityWithLatAndLon(25,14)

			if err != nil {
				assert.Error(t, err, test.expectedError.Error())
			} else {
				assert.Nil(t, test.expectedError)
			}
			assert.Equal(t, test.expectedFacility, facility)
		})

	}
}
