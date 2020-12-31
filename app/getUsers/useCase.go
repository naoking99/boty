package getusers

import (
	"github.com/naoking99/boty/infrastracturre/repository"
)

// UseCase acts the use case of sign-up.
func UseCase() *OutputData {
	users := repository.GetUsers()

	emails := []string{}

	for _, u := range *users {
		emails = append(emails, u.GetEmail())
	}

	return &OutputData{emails: emails}
}
