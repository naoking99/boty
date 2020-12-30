package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	signup "github.com/naoking99/boty/app/signUp"
)

// SignUp hundles sign-up's request and pass it to Usecase.
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Endpoint Hit: Sign Up!!")

		p := signup.CreateParam(r.FormValue("email"), r.FormValue("password"))
		u := signup.UseCase(p)

		json.NewEncoder(w).Encode(u)
	}
}
