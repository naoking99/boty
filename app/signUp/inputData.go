package signup

// InputData is
type InputData struct {
	email    string
	password string
}

// CreateInputData create parameters for the use case.
func CreateInputData(email string, password string) *InputData {
	return &InputData{email: email, password: password}
}
