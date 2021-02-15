package services

import (
	"musubi/src/context"
	maildir "musubi/src/services/maildir"
	"musubi/src/services/mailparser"
)

// GetMaildirService ...
func GetMaildirService() maildir.Service {
	return maildir.NewMaildirService(
		context.GetConfig(),
		context.GetLogger("maildir.Service"),
	)
}

// GetMailParserSvc
func GetMailParserSvc() mailparser.Service {
	return mailparser.NewMailParser(
		context.GetLogger("mailparser.Service"),
	)
}
