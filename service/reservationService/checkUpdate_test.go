package reservationService

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"spotHeroProject/lib"
	"spotHeroProject/mocks"
	"spotHeroProject/models"
	"testing"
)

func TestCheckUpdate(t *testing.T) {
	item := map[string]*dynamodb.AttributeValue{
		"reservation_id":&dynamodb.AttributeValue{
			S: aws.String("2"),
		},
		"update_possible":&dynamodb.AttributeValue{
			BOOL: aws.Bool(false),
		},
	}
	tests := []struct{
		name string
		errorGetItem error
		resultGetItem *dynamodb.GetItemOutput
		expectedError error
		expectedReservation models.Reservation
	}{
		{name: "getItemError", errorGetItem: errors.New("getItem thrown an error"),expectedError: errors.New("getItem thrown an error"),expectedReservation: models.Reservation{}},
		{name: "not found", resultGetItem: &dynamodb.GetItemOutput{},expectedError: lib.ErrNotFound,expectedReservation: models.Reservation{}},
		{name: "ok", resultGetItem: &dynamodb.GetItemOutput{Item: item},expectedReservation: models.Reservation{Id: "2",UpdatePossible: false}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db:=new(mocks.DynamoDBAPI)
			service:= New(db)
			db.On("GetItem",mock.Anything).Return(test.resultGetItem,test.errorGetItem)

			result, err := service.CheckUpdate("")


			if err != nil {
				assert.Contains(t, err.Error(), test.expectedError.Error())
			}else{
				assert.Nil(t, test.errorGetItem)
			}

			assert.Equal(t, test.expectedReservation,result)

		})
	}

}