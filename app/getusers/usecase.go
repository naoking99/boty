package getusers

import (
	"github.com/naoking99/boty/infrastracturre/repository/users"
)

// UseCase acts the use case of sign-up.
func UseCase() *OutputData {
	users := users.GetAll()

	emails := []string{}

	for _, u := range *users {
		emails = append(emails, u.Email())
	}

	return &OutputData{emails: emails}
}
