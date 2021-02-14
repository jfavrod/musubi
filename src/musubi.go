package main

import (
	"fmt"
	"musubi/src/context"
	"musubi/src/services"
	"musubi/src/services/maildir"
)

var (
	logger     context.Logger
	maildirsvc maildir.Service
)

func init() {
	logger = context.GetLogger("main")
	maildirsvc = services.GetMaildirService()
}

func main() {
	logger.Info("Starting...")
	emails := maildirsvc.GetEmailsInBlocked()
	fmt.Println(len(emails))
}
