package lib

import (
	"log"
	"os"
)

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
	if AWS_REGION == ""{
		log.Fatalln("AWS_REGION not found in environment variable ")
	}
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	if ACCESS_TOKEN == ""{
		log.Fatalln("ACCESS_TOKEN not found in environment variable ")
	}
	SECRET_KEY = os.Getenv("SECRET_KEY")
	if SECRET_KEY == ""{
		log.Fatalln("SECRET_KEY not found in environment variable ")
	}
}
