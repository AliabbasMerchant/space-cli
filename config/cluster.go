package config

import (
	"errors"

	"github.com/spaceuptech/space-cli/utils"
)

// AddCluster adds a cluster to the config
func AddCluster(name, url string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster Name Cannot be Empty")
	}
	if len(url) == 0 {
		return errors.New("Cluster URL Cannot be Empty")
	}

	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	_, ok := conf.Clusters[name]
	if !ok {
		conf.Clusters[name] = url
		return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
	}
	return errors.New(name + " already exists")
}

// RemoveCluster removes a cluster from the config
func RemoveCluster(name string) error {

	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster Name Cannot be Empty")
	}

	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	_, ok := conf.Clusters[name]
	if !ok {
		return errors.New(name + " does not exist")
	}

	delete(conf.Clusters, name)
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}

// SetClusterURL sets a cluster url
func SetClusterURL(name, url string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster Name Cannot be Empty")
	}
	if len(url) == 0 {
		return errors.New("Cluster URL Cannot be Empty")
	}

	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	_, ok := conf.Clusters[name]
	if !ok {
		return errors.New(name + " does not exist")
	}
	conf.Clusters[name] = url
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}
