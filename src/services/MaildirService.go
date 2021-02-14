package services

import (
	"io/ioutil"
	"musubi/src/context"
)

// MaildirService for working with a users maildir.
type MaildirService struct {
	config context.Config
	logger context.Logger
}

// New a new MaildirService object.
func newMaildirService(config context.Config, logger context.Logger) MaildirService {
	var svc MaildirService
	svc.config = config
	svc.logger = logger

	return svc
}

// GetEmailsInBlocked ...
func (svc MaildirService) GetEmailsInBlocked() []string {
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
