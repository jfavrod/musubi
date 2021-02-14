package services

import (
	"musubi/src/context"
	maildir "musubi/src/services/maildir"
)

// GetMaildirService ...
func GetMaildirService() maildir.Service {
	return maildir.NewMaildirService(
		context.GetConfig(),
		context.GetLogger("maildir.Service"),
	)
}
