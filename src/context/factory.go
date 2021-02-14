package context

// GetConfig get an new logger object.
func GetConfig() Config {
	return newConfig()
}

// GetLogger get an new logger object.
func GetLogger() Logger {
	return newLogger(newConfig())
}
