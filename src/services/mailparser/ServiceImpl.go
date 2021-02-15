package mailparser

import (
	"bufio"
	"musubi/src/context"
	"os"
	"regexp"
)

type mailparser struct {
	logger context.Logger
}

// NewMailParser ...
func NewMailParser(logger context.Logger) Service {
	var mp mailparser
	mp.logger = logger
	return mp
}

// GetSender ...
func (mp mailparser) GetSender(filepath string) Sender {
	var scanner *bufio.Scanner
	var sender Sender

	file, err := os.Open(filepath)

	if err != nil {
		mp.logger.Warn("Failed to read email file: " + filepath + "; " + err.Error())
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		line := []byte(scanner.Text())
		match, err := regexp.Match("Received: from", line)

		if err == nil && match {
			sender.ip = getIP(line)
		} else if err != nil {
			mp.logger.Warn("Failed to perform regex.Match on line: " + err.Error())
		}
	}

	return sender
}

func getIP(recpt []byte) string {
	ipPat, _ := regexp.Compile("(\\d+.){3}\\d+")
	return string(ipPat.Find(recpt))
}
