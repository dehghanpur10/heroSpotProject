package models

type Facility struct {
	Id        string `json:"id"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Latitude  int64  `json:"latitude"`
	Longitude int64  `json:"longitude"`
}
