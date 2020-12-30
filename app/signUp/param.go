package signup

// Param is
type Param struct {
	email    string
	password string
}

// CreateParam create parameters for the use case.
func CreateParam(email string, password string) *Param {
	return &Param{email: email, password: password}
}
