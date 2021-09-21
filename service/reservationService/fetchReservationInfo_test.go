package reservationService

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"spotHeroProject/lib"
	"spotHeroProject/mocks"
	"spotHeroProject/models"
	"testing"
)

func TestFetchReservationInfo(t *testing.T) {

	input := models.InputReservation{
		Id:             "1",
		UpdatePossible: false,
		Facility:       "2",
		ParkedVehicle:  "2",
	}
	output := models.Reservation{
		Id:             "1",
		UpdatePossible: false,
		ParkedVehicle: models.Vehicle{
			Id: "2",
		},
		Facility: models.Facility{
			Id:        "2",
			Latitude:  25,
			Longitude: 14,
		},
	}
	vehicleGetItemResult := map[string]*dynamodb.AttributeValue{
		"vehicle_id": {
			S: aws.String("2"),
		},
	}
	facilityGetItemResult := map[string]*dynamodb.AttributeValue{
		"facility_id": {
			S: aws.String("2"),
		},
		"latitude": {
			N: aws.String("25"),
		},
		"longitude": {
			N: aws.String("14"),
		},
	}
	tests := []struct {
		name                  string
		vehicleGetItemError   error
		facilityGetItemError  error
		vehicleGetItemResult  *dynamodb.GetItemOutput
		facilityGetItemResult *dynamodb.GetItemOutput
		expectedError         error
		reservation           models.Reservation
	}{
		{name: "ok", vehicleGetItemResult: &dynamodb.GetItemOutput{Item: vehicleGetItemResult}, facilityGetItemResult: &dynamodb.GetItemOutput{Item: facilityGetItemResult}, reservation: output},
		{name: "getVehicle GetItem error", vehicleGetItemError: errors.New("GetItem thrown an error"), expectedError: errors.New("GetItem thrown an error"), reservation: models.Reservation{}},
		{name: "getVehicle notFound error", vehicleGetItemResult: &dynamodb.GetItemOutput{}, expectedError: lib.ErrNotFound, reservation: models.Reservation{}},
		{name: "getFacility GetItem error", vehicleGetItemResult: &dynamodb.GetItemOutput{Item: vehicleGetItemResult}, expectedError: errors.New("GetItem thrown an error"), facilityGetItemError: errors.New("GetItem thrown an error"), reservation: models.Reservation{}},
		{name: "getFacility notFound error0", vehicleGetItemResult: &dynamodb.GetItemOutput{Item: vehicleGetItemResult}, facilityGetItemResult: &dynamodb.GetItemOutput{}, expectedError: lib.ErrNotFound, reservation: models.Reservation{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			db := new(mocks.DynamoDBAPI)
			db.On("GetItem", &dynamodb.GetItemInput{
				TableName: aws.String("VehicleSpot"),
				Key: map[string]*dynamodb.AttributeValue{
					"vehicle_id": &dynamodb.AttributeValue{
						S: aws.String("2"),
					},
				},
			}).Return(test.vehicleGetItemResult, test.vehicleGetItemError)
			db.On("GetItem", &dynamodb.GetItemInput{
				TableName: aws.String("FacilitySpot"),
				Key: map[string]*dynamodb.AttributeValue{
					"facility_id": &dynamodb.AttributeValue{
						S: aws.String("2"),
					},
				},
			}).Return(test.facilityGetItemResult, test.facilityGetItemError)
			service := New(db)
			// Act
			reservation, err := service.FetchReservationInfo(input)
			// Assert
			if err != nil {
				assert.Contains(t, err.Error(), test.expectedError.Error())
			} else {
				assert.Nil(t, test.expectedError)
			}
			assert.Equal(t, test.reservation, reservation)

		})
	}
}
