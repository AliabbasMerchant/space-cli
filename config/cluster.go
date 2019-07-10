package config

import (
  "errors"

  "github.com/spaceuptech/space-cli/utils"
)

// AddCluster adds a cluster to the config
func (conf *Config) AddCluster(name, url string) error {
	_, ok := conf.Clusters[name]
	if !ok {
		conf.Clusters[name] = url
		return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
	}
	return errors.New(name + " already exists")
}

// RemoveCluster removes a cluster from the config
func (conf *Config) RemoveCluster(name string) error {
	_, ok := conf.Clusters[name]
	if !ok {
		return errors.New(name + " does not exist")
	}
	delete(conf.Clusters, name)
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}

// SetClusterURL sets a cluster url
func (conf *Config) SetClusterURL(name, url string) error {
	_, ok := conf.Clusters[name]
	if !ok {
		return errors.New(name + " does not exist")
	}
	conf.Clusters[name] = url
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}
