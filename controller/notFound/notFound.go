package notFound

import (
	"net/http"
	"spotHeroProject/lib"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	lib.HttpError404(w, "Requested resource doesn't exist. Please check your path.")
}
