package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"spotHeroProject/controller/notFound"
	"spotHeroProject/controller/reservation"
	"spotHeroProject/controller/search"
	"spotHeroProject/controller/vehicle"
)

func Init() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v2/vehicles", vehicle.Create).Methods(http.MethodPost)
	r.HandleFunc("/v2/search", search.Search).Methods(http.MethodGet)
	r.HandleFunc("/v2/reservations", reservation.Create).Methods(http.MethodPost)
	r.HandleFunc("/v2/reservations", reservation.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/v2/reservations/{reservation_id}/update", reservation.CheckUpdate).Methods(http.MethodGet)
	r.HandleFunc("/v2/reservations/{reservation_id}/update", reservation.Update).Methods(http.MethodPut)
	r.NotFoundHandler = http.HandlerFunc(notFound.Handler)

	return r
}
