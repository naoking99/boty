package users

import (
	"github.com/naoking99/boty/domain/user"
)

// Repository is
type Repository interface {
	GetAll() *[]*user.User
	Save(u *user.User)
}
