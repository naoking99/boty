package controller

import (
	"encoding/json"
	"net/http"

	"github.com/naoking99/boty/app/getusers"
	"github.com/naoking99/boty/infrastracture/users"
)

// User hundles sign-up's request and pass it to Usecase.
func User(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		o := getusers.UseCase(users.Repository{})
		json.NewEncoder(w).Encode(o.Emails())
	}
}
