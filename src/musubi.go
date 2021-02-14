package main

import (
	"fmt"
	"io/ioutil"
	"musubi/src/config"
	"musubi/src/logging"
)

var conf = config.New()
var log = logging.New(conf)

func main() {
	log.Info.Print("Starting...")
	emails := getEmailsInBlocked()
	fmt.Println(len(emails))
}

func getEmailsInBlocked() []string {
	var emails = make([]string, 5)
	fileInfo, err := ioutil.ReadDir(conf.BlockDir)

	if err != nil {
		log.Error.Print("Failed to read directory: ", conf.BlockDir)
		return emails
	}

	for _, entry := range fileInfo {
		if entry.IsDir() != true {
			emails = append(emails, entry.Name())
		}
	}

	return emails
}
