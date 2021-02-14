package main

import (
	"fmt"
	"musubi/src/context"
	"musubi/src/services"
)

var logger = context.GetLogger()
var maildirSvc = services.GetMaildirService()

func main() {
	logger.Info("Starting...")
	emails := maildirSvc.GetEmailsInBlocked()
	fmt.Println(len(emails))
}
