package controller

import (
	"fmt"
	"net/http"
)

func GetHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
	fmt.Println("Endpoint Hit: homePage")
}
