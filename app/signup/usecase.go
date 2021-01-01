package signup

import (
	"context"
	"log"

	"firebase.google.com/go/v4/auth"
	"github.com/naoking99/boty/app/repository/users"
	"github.com/naoking99/boty/domain/user"
	"github.com/naoking99/boty/utils/firebase"
)

// UseCase acts the use case of sign-up.
func UseCase(i *Input, usersRepo users.Repository) *Output {
	params := (&auth.UserToCreate{}).
		Email(i.email).
		EmailVerified(false).
		Password(i.password)

	_, err := firebase.AuthClient().CreateUser(context.Background(), params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}

	u := user.New(0, "testname", i.email)
	usersRepo.Save(u)

	return &Output{email: i.email}
}
