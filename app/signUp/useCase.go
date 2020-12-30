package signup

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// UseCase acts the use case of sign-up.
func UseCase(p *Param) *auth.UserRecord {
	fmt.Printf("email is %s, password is %s in UseCase", p.email, p.password)

	ctx := context.Background()
	opt := option.WithCredentialsFile("boty-dev-firebase-adminsdk.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	params := (&auth.UserToCreate{}).
		Email(p.email).
		EmailVerified(false).
		Password(p.password)

	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)

	return u
}
