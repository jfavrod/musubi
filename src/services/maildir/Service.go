package maildir

// Service ...
type Service interface {
	GetEmailsInBlocked() []string
}
