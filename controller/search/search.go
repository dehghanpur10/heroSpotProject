package search

import (
	"fmt"
	"net/http"
	"spotHeroProject/lib"
)

// SearchFacilityController
// @Summary search facility on based their lat and lon
// @Description this endpoint will search facility on based their lan and lon
// @Tags search
// @Accept  json
// @Produce  json
// @Param lat query string true "Latitude"
// @Param lon query string true "longitude"
// @Success 200 {array} models.Facility
// @Failure 400 {object} lib.ErrorResponse
// @Failure 500 {object} lib.ErrorResponse
// @Router /v2/search [Get]
func SearchFacilityController(w http.ResponseWriter, r *http.Request) {
	lib.InitLog(r)

	lat := r.FormValue("lat")
	lon := r.FormValue("lon")
	if lon == "" || lat == "" {
		fmt.Println("SearchFacilityController -  bad request")
		lib.HttpError400(w, "lan and lon should be send in query params")
		return
	}

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Printf("SearchFacilityController - %v", err)
		lib.HttpError500(w)
		return
	}
	_ = db

	// Todo add search service

}
