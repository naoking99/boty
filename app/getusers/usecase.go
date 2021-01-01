package getusers

import (
	"github.com/naoking99/boty/app/repository/users"
)

// UseCase acts the use case of sign-up.
func UseCase(usersRepo users.Repository) *Output {
	users := usersRepo.GetAll()

	emails := []string{}

	for _, u := range *users {
		emails = append(emails, u.Email())
	}

	return &Output{emails: emails}
}
