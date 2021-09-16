package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func InitLog(r *http.Request) {
	fmt.Printf("%s-%s-%v-", r.Method, r.URL.Path, r.URL.Query())
}

func jsonResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
}

func HttpSuccessResponse(w http.ResponseWriter, statusCode int, payload []byte) {
	jsonResponseHeaders(w)
	w.WriteHeader(statusCode)
	if statusCode == http.StatusNoContent {
		return
	}
	_, err := w.Write(payload)
	if err != nil {
		fmt.Print("HttpSuccessResponse-Write Error:", err)
		HttpError500(w)
	}
}

func HttpErrorResponse(w http.ResponseWriter, statusCode int, title, description string) {
	jsonResponseHeaders(w)
	w.WriteHeader(statusCode)
	errorResponse := NewErrorResponse(statusCode, title, description)
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		fmt.Print("HttpErrorResponse-json.Encode Error:", err)
		HttpError500(w)
	}
}

func HttpError400(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusBadRequest, "Bad request", description)
}

func HttpError401(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusUnauthorized, "Unauthorized", description)
}

func HttpError402(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusPaymentRequired, "Payment Required", description)
}

func HttpError403(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusForbidden, "Forbidden", description)
}

func HttpError404(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusNotFound, "Not found", description)
}

func HttpError500(w http.ResponseWriter) {
	HttpErrorResponse(w, http.StatusInternalServerError, "Internal error", "Internal server error.")
}

func HttpError503(w http.ResponseWriter) {
	HttpErrorResponse(w, http.StatusServiceUnavailable, "Service Unavailable", "Service is unavailable.")
}

func HttpErrorWith(w http.ResponseWriter, err error) {
	if err == nil {
		fmt.Print("HttpErrorWith: Error was nil.")
		HttpError500(w)
		return
	}
	if errors.Is(err, ErrNotFound) {
		HttpError404(w, err.Error())
	} else if errors.Is(err, ErrForbidden) {
		HttpError403(w, err.Error())
	} else if errors.Is(err, ErrDatabaseConnection) {
		HttpError500(w)
	} else {
		HttpError500(w)
	}
}
