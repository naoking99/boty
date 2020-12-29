package signup

// Param is
type Param struct {
	Email    string
	Password string
}

// CreateParam create parameters for the use case.
func CreateParam(email string, password string) *Param {
	return &Param{Email: email, Password: password}
}
