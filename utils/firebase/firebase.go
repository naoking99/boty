package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var app = initFirebaseApp()
var client = initAuth()

func initFirebaseApp() *firebase.App {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func initAuth() *auth.Client {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client
}

// AuthClient returns an instance of auth.Client.
func AuthClient() *auth.Client {
	return client
}
