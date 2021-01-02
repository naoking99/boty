package router

import (
	"log"
	"net/http"

	"github.com/naoking99/boty/presentation/controller"
)

// HandleRequests handles http requests.
func HandleRequests() {
	http.HandleFunc("/sign-up", controller.SignUp)
	http.HandleFunc("/users", controller.User)
	http.HandleFunc("/", controller.HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
