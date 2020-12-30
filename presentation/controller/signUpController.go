package controller

import (
	"fmt"
	"net/http"

	signup "github.com/naoking99/boty/app/signUp"
)

// SignUp hundles sign-up's request and pass it to Usecase.
func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// p := signup.UseCase(signup.CreateParam(r.FormValue("email"), r.FormValue("password")))
		p := signup.CreateParam(r.FormValue("email"), r.FormValue("password"))
		signup.UseCase(p)
		fmt.Println("Endpoint Hit: Sign Up!!")
	}
}
