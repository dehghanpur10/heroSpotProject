package reservationService

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"spotHeroProject/lib"
	"spotHeroProject/models"
)

func (s *ReservationService) FetchReservationInfo(reservation models.InputReservation) (models.Reservation, error) {
	var completeReservation models.Reservation

	vehicleId := reservation.ParkedVehicle
	facilityId := reservation.Facility

	err := s.getVehicle(&completeReservation, vehicleId)
	if err != nil {
		err = fmt.Errorf("reservationService.FetchReservationInfo  - %v ", err)
		return models.Reservation{}, err
	}

	err = s.getFacility(&completeReservation, facilityId)
	if err != nil {
		err = fmt.Errorf("reservationService.FetchReservationInfo - %v ", err)
		return models.Reservation{}, err
	}

	return completeReservation, nil

}

func (s *ReservationService) getVehicle(reservation *models.Reservation, vehicleId string) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Vehicle"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(vehicleId),
			},
		},
	}

	result, err := s.db.GetItem(input)
	if err != nil {
		err = fmt.Errorf("getVehicle.GetItem - %v - %v", lib.ErrInternal, err)
		return err
	}

	if result.Item == nil {
		err = fmt.Errorf("getVehicle - %v", lib.ErrNotFound)
		return err
	}

	var vehicle models.Vehicle
	err = dynamodbattribute.UnmarshalMap(result.Item, &vehicle)
	if err != nil {
		err = fmt.Errorf("getVehicle.unmarshalMap - %v - %v", lib.ErrInternal, err)
		return err
	}

	reservation.ParkedVehicle = vehicle
	return nil
}

func (s *ReservationService) getFacility(reservation *models.Reservation, facilityId string) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Facility"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(facilityId),
			},
		},
	}

	result, err := s.db.GetItem(input)
	if err != nil {
		err = fmt.Errorf("getFacility.GetItem - %v - %v", lib.ErrInternal, err)
		return err
	}

	if result.Item == nil {
		err = fmt.Errorf("getFacility - %v", lib.ErrNotFound)
		return err
	}

	var facility models.Facility
	err = dynamodbattribute.UnmarshalMap(result.Item, &facility)
	if err != nil {
		err = fmt.Errorf("getFacility.unmarshalMap - %v - %v", lib.ErrInternal, err)
		return err
	}

	reservation.Facility = facility
	return nil
}
