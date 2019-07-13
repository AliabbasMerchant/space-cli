package config

import (
	"errors"

	"github.com/spaceuptech/space-cli/utils"
)

// RemoveEnvVar removes an environment variable
func RemoveEnvVar(name string) error {

	// Sanity check
	if len(name) == 0 {
		return errors.New("Environment variable name cannot be empty")
	}

	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	_, ok := conf.Env[name]
	if !ok {
		return errors.New(name + " does not exist")
	}

	delete(conf.Env, name)
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}

// SetEnvVar sets an environment variable
func SetEnvVar(name, value string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Environment variable name cannot be empty")
	}
	if len(value) == 0 {
		return errors.New("Environment variable value cannot be empty")
	}

	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	conf.Env[name] = value
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}
