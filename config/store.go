package config

import (
  "errors"
  "io/ioutil"
	"strings"
	"encoding/json"

  "gopkg.in/yaml.v2"
)

// StoreDeployerToFile stores the deployer file to disk
func StoreConfigToFile(conf *Deploy, path string) error {
	var data []byte
	var err error

	if strings.HasSuffix(path, ".yaml") {
		data, err = yaml.Marshal(conf)
	} else {
		return errors.New("Invalid config file type")
	}

	// Check if error occured while marshaling
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, data, 0644)
}

// GetJSON gets the json bytes of the config
func GetJSON(conf *Deploy) ([]byte, error) {
	return json.Marshal(conf)
}
