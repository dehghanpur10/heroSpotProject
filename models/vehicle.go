package models

type Vehicle struct {
	Id          string             `json:"vehicle_id" validate:"required" example:"1"`
	Description VehicleDescription `json:"vehicle_description" validate:"required"`
}

type VehicleDescription struct {
	Name  string `json:"name" validate:"required" example:"benz"`
	Model string `json:"model" validate:"required" example:"s300"`
	Year  string `json:"year" validate:"required" example:"2021"`
}
