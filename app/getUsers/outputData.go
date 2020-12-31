package getusers

// OutputData is
type OutputData struct {
	emails []string
}

// GetEmails is
func (o OutputData) GetEmails() []string {
	return o.emails
}
