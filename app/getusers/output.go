package getusers

// Output is
type Output struct {
	emails []string
}

// Emails is
func (o Output) Emails() []string {
	return o.emails
}
