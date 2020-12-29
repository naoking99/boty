package controller

import (
	"fmt"
	"net/http"
)

// GetHomePage returns Home page's JSON.
func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homePage")
}
