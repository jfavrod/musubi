package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config for musibi.
type Config struct {
	// BlockDir the directory where user puts emails to be blocked.
	BlockDir   string
	BlockList  string
	BlockedDir string
	LogFile    string
	MailDir    string
}

// New config object.
func New() Config {
	var config Config
	var configFile = getConfigFile()

	json.Unmarshal(configFile, &config)

	config.BlockDir = config.MailDir + "/" + config.BlockDir
	config.BlockedDir = config.MailDir + "/" + config.BlockedDir
	config.BlockList = config.MailDir + "/" + config.BlockList

	return config
}

func getConfigDir() string {
	var dir = ""
	var env = os.Getenv("GO_ENV")

	if env == "DEV" {
		dir = "/home/jason/dev/musubi/dev/etc"
	}

	return dir
}

func getConfigFile() []byte {
	var dir = getConfigDir()

	data, err := ioutil.ReadFile(dir + "/musubi.json")

	if err == nil {
		return data
	}

	panic(err)
}
