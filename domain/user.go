package domain

// User is
type User struct {
	id    int
	name  string
	email string
}

// CreateUser is
func CreateUser(id int, name string, email string) *User {
	return &User{id: id, name: name, email: email}
}

// GetEmail is
func (u *User) GetEmail() string {
	return u.email
}
