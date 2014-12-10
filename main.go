package configo

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	Title string
	Port  string
}

// Reads info from config file
func ReadConfig(confFile string) Config {

	f, err := os.Open(confFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}

	return config
}
