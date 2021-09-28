package models

type Facility struct {
	Id        string `json:"facility_id"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}
