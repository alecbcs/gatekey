package config

import (
	"flag"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type config struct {
	General        general
	Authentication authentication
	Database       database
	Relay          relay
	Tokens         tokens
}

type general struct {
	Version string
	Port    int
}
type authentication struct {
	User     string
	Password string
}

type database struct {
	Location string
}

type relay struct {
	Location         string
	User             string
	Password         string
	TempFileLocation string
}

type tokens struct {
	Length int
}

var (
	// Global is the configuration struct for the application.
	Global config
	path   string
)

// initialize the app config system. If a config doesn't exist, create one.
// If the config is out of date read the current config and rebuild with new fields.
func init() {
	// Determine the current user to build expected file path.
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// Create expected config path.
	flagStr := flag.String("config", filepath.Join(user.HomeDir, ".config", "gatekey", "gatekey.config"), "string")
	flag.Parse()
	path = *(flagStr)

	Global = defaultConf()
	readConf(&Global)
	// If the configuration version has changed update the config to the new
	// format while keeping the user's preferences.
	if Global.General.Version != defaultConf().General.Version {
		reloadConf()
		readConf(&Global)
	}
}

// Read the config or create a new one if it doesn't exist.
func readConf(conf *config) {
	_, err := toml.DecodeFile(path, &conf)
	if os.IsNotExist(err) {
		genConf(defaultConf())
		readConf(conf)
	}
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
}
