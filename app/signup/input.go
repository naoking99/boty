package signup

// Input is
type Input struct {
	email    string
	password string
}

// NewInput create parameters for the use case.
func NewInput(email string, password string) *Input {
	return &Input{email: email, password: password}
}
