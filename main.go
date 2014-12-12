// Package configo provides some functions to handle configuration
// file in TOML format
package configo

import (
	"github.com/naoina/toml"
	"io/ioutil"
	"os"
)

// ReadConfig reads from config file
func ReadConfig(configFile string, o interface{}) (err error) {

	f, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(buf, o); err != nil {
		panic(err)
	}

	return err

}
