package maildir

import (
	"io/ioutil"
	"musubi/src/context"
)

// MaildirService for working with a users maildir.
type maildirService struct {
	config context.Config
	logger context.Logger
}

// NewMaildirService a new MaildirService object.
func NewMaildirService(config context.Config, logger context.Logger) Service {
	var svc maildirService
	svc.config = config
	svc.logger = logger

	return svc
}

// GetEmailsInBlocked ...
func (svc maildirService) GetEmailsInBlocked() []string {
	var emails = make([]string, 5)
	fileInfo, err := ioutil.ReadDir(svc.config.BlockDir)

	if err != nil {
		svc.logger.Error("Failed to read directory: " + svc.config.BlockDir)
		return emails
	}

	for _, entry := range fileInfo {
		if entry.IsDir() != true {
			emails = append(emails, entry.Name())
		}
	}

	return emails
}
