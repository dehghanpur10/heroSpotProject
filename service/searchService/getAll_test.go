package searchService

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

func TestGetAllFacility(t *testing.T) {

	item := []map[string]*dynamodb.AttributeValue{
		{
			"facility_id": {
				S: aws.String("1"),
			},
			"city": {
				S: aws.String("Tehran"),
			},
			"country": {
				S: aws.String("Iran"),
			},
			"latitude": {
				N: aws.String("25"),
			},
			"longitude": {
				N: aws.String("14"),
			},
		},
	}
	facilities := []models.Facility{
		{Id: "1", City: "Tehran", Country: "Iran", Latitude: 25, Longitude: 14},
	}
	tests := []struct {
		name             string
		scanError        error
		scanOutput       *dynamodb.ScanOutput
		expectedFacility []models.Facility
		expectedError    error
	}{
		{name: "ok", scanOutput: &dynamodb.ScanOutput{Items: item}, expectedFacility: facilities},
		{name: "scan error", scanError: errors.New("scan error"), expectedError: errors.New("scan error")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			db := new(mocks.DynamoDBAPI)
			service := NewSearchService(db)
			db.On("Scan", mock.Anything).Return(test.scanOutput, test.scanError)
			// Act
			facility, err := service.GetAllFacility()
			// Assert
			if err != nil {
				assert.Error(t, err, test.expectedError.Error())
			} else {
				assert.Nil(t, test.expectedError)
			}
			assert.Equal(t, test.expectedFacility, facility)
		})

	}
}
