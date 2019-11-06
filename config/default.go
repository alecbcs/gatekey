package config

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

func defaultConf() config {
	result := config{
		General: general{
			Version: "0.0.1",
			Port:    8080,
		},
		Authentication: authentication{
			User:     "",
			Password: "",
		},
		Database: database{
			Location: filepath.Join(filepath.Dir(path), "tokens.db"),
		},
		Relay: relay{
			Location:         "",
			User:             "",
			Password:         "",
			TempFileLocation: filepath.Join(filepath.Dir(path), "temp"),
		},
		Tokens: tokens{
			Length: 32,
		},
	}
	return result
}

// genConf encodes the values of the Config stuct back into a TOML file.
func genConf(conf config) {
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(conf)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(buf.Bytes())
}

// reloadConf imports the users config onto a default config and then rewrites
// the configuration file.
func reloadConf() {
	result := defaultConf()
	readConf(&result)
	result.General.Version = defaultConf().General.Version
	genConf(result)
}
