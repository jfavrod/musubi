package services

import (
	"musubi/src/context"
)

// GetMaildirService ...
func GetMaildirService() MaildirService {
	return newMaildirService(
		context.GetConfig(),
		context.GetLogger(),
	)
}
