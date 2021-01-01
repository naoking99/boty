package controller

import (
	"encoding/json"
	"net/http"

	"github.com/naoking99/boty/app/signup"
	"github.com/naoking99/boty/infrastracture/users"
)

// SignUp hundles sign-up's request and pass it to Usecase.
func SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		i := signup.NewInput(r.FormValue("email"), r.FormValue("password"))
		o := signup.UseCase(i, users.Repository{})

		res := map[string]string{
			"email": o.Email(),
		}

		json.NewEncoder(w).Encode(res)
	}
}
