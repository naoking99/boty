package getusers

// OutputData is
type OutputData struct {
	emails []string
}

// Emails is
func (o OutputData) Emails() []string {
	return o.emails
}
