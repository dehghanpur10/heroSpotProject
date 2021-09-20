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

var item = []map[string]*dynamodb.AttributeValue{
	{
		"facility_id": &dynamodb.AttributeValue{
			S: aws.String("1"),
		},
		"city": &dynamodb.AttributeValue{
			S: aws.String("Tehran"),
		},
		"country": &dynamodb.AttributeValue{
			S: aws.String("Iran"),
		},
		"latitude": &dynamodb.AttributeValue{
			N: aws.String("25"),
		},
		"longitude": &dynamodb.AttributeValue{
			N: aws.String("14"),
		},
	},
}
var facilities = []models.Facility{
	{Id: "1", City: "Tehran", Country: "Iran", Latitude: 25, Longitude: 14},
}

func TestGetAllFacility(t *testing.T) {
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
			db := new(mocks.DynamoDBAPI)
			service := NewSearchService(db)
			db.On("Scan", mock.Anything).Return(test.scanOutput, test.scanError)

			facility, err := service.GetAllFacility()
			if err != nil {
				assert.Error(t, err, test.expectedError.Error())
			} else {
				assert.Nil(t, test.expectedError)
			}
			assert.Equal(t, test.expectedFacility, facility)
		})

	}
}
