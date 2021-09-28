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
	completeReservation.Quote = reservation.Quote
	completeReservation.UpdatePossible = reservation.UpdatePossible
	completeReservation.Id = reservation.Id

	vehicleId := reservation.ParkedVehicle
	facilityId := reservation.Facility

	err := s.getVehicle(&completeReservation, vehicleId)
	if err != nil {
		fmt.Print("reservationService.FetchReservationInfo  - ")
		err = fmt.Errorf("%w ", err)
		return models.Reservation{}, err
	}

	err = s.getFacility(&completeReservation, facilityId)
	if err != nil {
		fmt.Print("reservationService.FetchReservationInfo - ")
		err = fmt.Errorf("%w ", err)
		return models.Reservation{}, err
	}

	return completeReservation, nil
}

func (s *ReservationService) getVehicle(reservation *models.Reservation, vehicleId string) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(lib.VEHICLE_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"vehicle_id": {
				S: aws.String(vehicleId),
			},
		},
	}

	result, err := s.db.GetItem(input)
	if err != nil {
		fmt.Print("getVehicle.GetItem - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return err
	}

	if result.Item == nil {
		fmt.Print("getVehicle - ")
		err = fmt.Errorf("%w", lib.ErrNotFound)
		return err
	}

	var vehicle models.Vehicle
	err = dynamodbattribute.UnmarshalMap(result.Item, &vehicle)
	if err != nil {
		fmt.Print("getVehicle.unmarshalMap - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return err
	}

	reservation.ParkedVehicle = vehicle
	return nil
}

func (s *ReservationService) getFacility(reservation *models.Reservation, facilityId string) error {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(lib.FACILITY_TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"facility_id": {
				S: aws.String(facilityId),
			},
		},
	}

	result, err := s.db.GetItem(input)
	if err != nil {
		fmt.Print("getFacility.GetItem - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return err
	}

	if result.Item == nil {
		fmt.Print("getFacility - ")
		err = fmt.Errorf("%w", lib.ErrNotFound)
		return err
	}

	var facility models.Facility
	err = dynamodbattribute.UnmarshalMap(result.Item, &facility)
	if err != nil {
		fmt.Print("getFacility.unmarshalMap - ")
		err = fmt.Errorf("%w - %v", lib.ErrInternal, err)
		return err
	}

	reservation.Facility = facility
	return nil
}
