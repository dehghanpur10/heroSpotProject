package models

type Reservation struct {
	Id             string   `json:"reservation_id"`
	ParkedVehicle  Vehicle  `json:"parked_vehicle"`
	Quote          Quote    `json:"quote"`
	Facility       Facility `json:"facility"`
	UpdatePossible bool     `json:"update_possible"`
}

type InputReservation struct {
	Id             string `json:"reservation_id" validate:"required" example:"1"`
	ParkedVehicle  string `json:"parked_vehicle_id" validate:"required" example:"1"`
	Quote          Quote  `json:"quote"`
	Facility       string `json:"facility_id" validate:"required" example:"1"`
	UpdatePossible bool   `json:"update_possible" validate:"required" example:"true"`
}

type Quote struct {
	Starts string `json:"starts" validate:"required" example:"2019-08-19T13:49:37.000Z"`
	Ends   string `json:"ends" validate:"required" example:"2019-08-19T13:49:37.000Z"`
}
