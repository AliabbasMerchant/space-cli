package config

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/spaceuptech/space-cli/model"
)

// StoreConfigToFile stores the deployer file to disk
func StoreConfigToFile(conf *model.Deploy, path string) error {
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
