package controller

import (
	"fmt"
	"net/http"
)

// SignUp hundles sign-up's request and pass it to Usecase.
func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sign Up!!")
	fmt.Println("Endpoint Hit: Sign Up!!")
}
