package context

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

// NewConfig config object.
func newConfig() Config {
	var conf Config
	var configFile = getConfigFile()

	json.Unmarshal(configFile, &conf)

	conf.BlockDir = conf.MailDir + "/" + conf.BlockDir
	conf.BlockedDir = conf.MailDir + "/" + conf.BlockedDir
	conf.BlockList = conf.MailDir + "/" + conf.BlockList

	return conf
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
