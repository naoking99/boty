package controller

import (
	"fmt"
	"net/http"
)

// HomePage returns Home page's JSON.
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homePage")
}
