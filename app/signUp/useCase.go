package signup

import (
	"context"
	"log"

	"firebase.google.com/go/auth"
	"github.com/naoking99/boty/utils/firebase"
)

// UseCase acts the use case of sign-up.
func UseCase(p *Param) *auth.UserRecord {
	params := (&auth.UserToCreate{}).
		Email(p.email).
		EmailVerified(false).
		Password(p.password)

	u, err := firebase.AuthClient().CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	log.Printf("Successfully created user: %#v\n", u.UserInfo)

	return u
}
