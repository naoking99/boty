package controller

import (
	"encoding/json"
	"net/http"

	signup "github.com/naoking99/boty/app/signUp"
)

// SignUp hundles sign-up's request and pass it to Usecase.
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		p := signup.CreateParam(r.FormValue("email"), r.FormValue("password"))
		o := signup.UseCase(p)

		res := map[string]string{
			"email": o.Email(),
		}

		json.NewEncoder(w).Encode(res)
	}
}
