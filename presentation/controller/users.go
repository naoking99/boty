package controller

import (
	"encoding/json"
	"net/http"

	"github.com/naoking99/boty/app/getusers"
)

// User hundles sign-up's request and pass it to Usecase.
func User(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		o := getusers.UseCase()
		json.NewEncoder(w).Encode(o.Emails())
	}
}