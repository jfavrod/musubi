package main

import (
	"fmt"
	"musubi/src/context"
	"musubi/src/services"
	"musubi/src/services/maildir"
	"musubi/src/services/mailparser"
)

var (
	logger        context.Logger
	maildirSvc    maildir.Service
	mailparserSvc mailparser.Service
)

func init() {
	logger = context.GetLogger("main")
	maildirSvc = services.GetMaildirService()
	mailparserSvc = services.GetMailParserSvc()
}

func main() {
	var sender mailparser.Sender
	emails := maildirSvc.GetEmailsInBlocked()

	for _, email := range emails {
		sender = mailparserSvc.GetSender(email)
		fmt.Println(sender)
	}
}
