package user

// User is
type User struct {
	id    int
	name  string
	email string
}

// New is
func New(id int, name string, email string) *User {
	return &User{id: id, name: name, email: email}
}

// Email is
func (u *User) Email() string {
	return u.email
}
