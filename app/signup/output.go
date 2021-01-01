package signup

// Output is
type Output struct {
	email string
}

// Email is
func (o Output) Email() string {
	return o.email
}
