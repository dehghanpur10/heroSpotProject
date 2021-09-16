package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spotHeroProject/lib"
	"spotHeroProject/models"
	"spotHeroProject/service/searchService"
	"strconv"
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

	db, err := lib.GetDynamoDB()
	if err != nil {
		fmt.Printf("SearchFacilityController - %v", err)
		lib.HttpError500(w)
		return
	}
	var facilities []models.Facility
	service := searchService.New(db)

	if lon == "" && lat == "" {
		facilities, err = service.GetAllFacility()
	} else {
		latNumber, err := strconv.ParseFloat(lat, 32)
		if err != nil {
			fmt.Printf("SearchFacilityController - ParseFloat %v", err)
			lib.HttpError400(w, "lat query should be number type")
			return
		}
		lonNumber, err := strconv.ParseFloat(lon, 32)
		if err != nil {
			fmt.Printf("SearchFacilityController - ParseFloat %v", err)
			lib.HttpError400(w, "lon query should be number type")
			return
		}
		facilities, err = service.GetFacilityWithLatAndLon(latNumber, lonNumber)
	}

	if err != nil {
		fmt.Printf("SearchFacilityController - %v", err)
		lib.HttpError500(w)
		return
	}

	result, err := json.Marshal(facilities)
	if err != nil {
		fmt.Printf("SearchFacilityController - Marshal - %v", err)
		lib.HttpError500(w)
		return
	}

	lib.HttpSuccessResponse(w, http.StatusOK, result)
}
