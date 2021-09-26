package lib

import "os"

var (
	AWS_REGION   string
	ACCESS_TOKEN string
	SECRET_KEY   string

	FACILITY_TABLE_NAME    = "FacilitySpot"
	RESERVATION_TABLE_NAME = "ReservationSpot"
	VEHICLE_TABLE_NAME     = "VehicleSpot"
)

func init() {
	AWS_REGION = os.Getenv("AWS_REGION")
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	SECRET_KEY = os.Getenv("SECRET_KEY")
}
