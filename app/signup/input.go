package signup

// InputData is
type InputData struct {
	email    string
	password string
}

// NewInputData create parameters for the use case.
func NewInputData(email string, password string) *InputData {
	return &InputData{email: email, password: password}
}
