package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"github.com/naoking99/boty/presentation/routes"
)

func initializeAppDefault() *firebase.App {
	// [START initialize_app_default_golang]
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_default_golang]

	return app
}

func main() {
	routes.HandleRequests()
}
