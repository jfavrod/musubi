package logging

import (
	"log"
	"musubi/src/config"
	"os"
)

// Logger for logging to the configured logfile.
type Logger struct {
	Error *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
}

// New logger object.
func New(conf config.Config) Logger {
	var logger Logger
	logfile, err := os.OpenFile(conf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	logger.Error = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warn = log.New(logfile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
