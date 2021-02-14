package context

import (
	"log"
	"os"
)

// Logger for logging to the configured logfile.
type Logger struct {
	error *log.Logger
	info  *log.Logger
	warn  *log.Logger
}

func newLogger(conf Config) Logger {
	var logger Logger
	logfile, err := os.OpenFile(conf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	logger.error = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.warn = log.New(logfile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}

// Error print an error message to the configured log file.
func (logger Logger) Error(message string) {
	logger.error.Print(message)
}

// Info print an info message to the configured log file.
func (logger Logger) Info(message string) {
	logger.info.Print(message)
}

// Warn print a warning message to the configured log file.
func (logger Logger) Warn(message string) {
	logger.warn.Print(message)
}
