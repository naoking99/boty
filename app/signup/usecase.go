package signup

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/naoking99/boty/utils/firebase"
)

// UseCase acts the use case of sign-up.
func UseCase(i *Input) *OutputData {
	params := (&auth.UserToCreate{}).
		Email(i.email).
		EmailVerified(false).
		Password(i.password)

	u, err := firebase.AuthClient().CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}

	return &OutputData{email: u.Email}
}
