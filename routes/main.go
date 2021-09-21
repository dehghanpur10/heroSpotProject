package routes

import (
	"net/http"
	"spotHeroProject/controller/httpserver"
	"spotHeroProject/controller/reservation"
	"spotHeroProject/controller/search"
	"spotHeroProject/controller/vehicle"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v2/vehicles", vehicle.CreateVehicleController).Methods(http.MethodPost)
	r.HandleFunc("/v2/search", search.SearchFacilityController).Methods(http.MethodGet)
	r.HandleFunc("/v2/reservations", reservation.CreateReservationController).Methods(http.MethodPost)
	r.HandleFunc("/v2/reservations", reservation.GetAllReservationController).Methods(http.MethodGet)
	r.HandleFunc("/v2/reservations/{reservation_id}/update", reservation.CheckUpdateReservationController).Methods(http.MethodGet)
	r.HandleFunc("/v2/reservations/{reservation_id}/update", reservation.UpdateReservationController).Methods(http.MethodPut)
	r.NotFoundHandler = http.HandlerFunc(httpserver.NotFoundController)

	return r
}
