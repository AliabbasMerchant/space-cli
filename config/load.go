package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/spaceuptech/space-cli/model"
	"github.com/spaceuptech/space-cli/utils"
)

// LoadGlobalConfigFromFile loads the global config from file
func LoadGlobalConfigFromFile(path string) (*model.GlobalConfig, error) {
	c := &model.GlobalConfig{Clusters: model.Clusters{}}

	// Load the file in memory
	data, err := ioutil.ReadFile(path)
	if err != nil {
		os.MkdirAll(utils.GetGlobalConfigDir(), os.ModePerm)
		StoreGlobalConfigToFile(c, path)
		return c, nil
	}

	if err := json.Unmarshal(data, c); err != nil {
		return nil, err
	}

	return c, nil
}

// LoadConfigFromFile loads the config from the provided file path
func LoadConfigFromFile(path string) (*model.Deploy, error) {
	// Load the file in memory
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("No config file present")
	}
	if strings.HasSuffix(path, ".yaml") {
		return loadConfigFromYAML(data)
	} else if strings.HasSuffix(path, ".json") {
		return loadConfigFromJSON(data)
	}

	return nil, errors.New("Invalid config file type provided")
}

// LoadConfigFromYAML loads the config from the provided yaml bytes
func loadConfigFromYAML(text []byte) (*model.Deploy, error) {
	// Marshal the configuration
	conf := new(model.Deploy)
	err := yaml.Unmarshal(text, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// LoadConfigFromJSON loads the config from the provided yaml bytes
func loadConfigFromJSON(text []byte) (*model.Deploy, error) {
	// Marshal the configuration
	conf := new(model.Deploy)
	err := json.Unmarshal(text, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
