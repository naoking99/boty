package signup

// OutputData is
type OutputData struct {
	email string
}

// Email is
func (o OutputData) Email() string {
	return o.email
}
