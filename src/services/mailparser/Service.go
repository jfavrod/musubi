package mailparser

// Sender ...
type Sender struct {
	hostname string
	ip       string
	name     string
}

// Service ...
type Service interface {
	GetSender(filepath string) Sender
}
