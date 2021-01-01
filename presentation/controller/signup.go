package controller

import (
	"encoding/json"
	"net/http"

	"github.com/naoking99/boty/app/signup"
)

// SignUp hundles sign-up's request and pass it to Usecase.
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		i := signup.NewInput(r.FormValue("email"), r.FormValue("password"))
		o := signup.UseCase(i)

		res := map[string]string{
			"email": o.Email(),
		}

		json.NewEncoder(w).Encode(res)
	}
}
