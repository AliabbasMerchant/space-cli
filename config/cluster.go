package config

import (
	"errors"

	"github.com/spaceuptech/space-cli/model"
	"github.com/spaceuptech/space-cli/utils"
)

// AddCluster adds a cluster to the config
func AddCluster(name, url string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster name needs to be provided")
	}
	if len(url) == 0 {
		return errors.New("Cluster url needs to be provided")
	}

	// Load config from file
	conf, err := LoadGlobalConfigFromFile(utils.GetGlobalConfigFile())
	if err != nil {
		return err
	}

	_, ok := conf.Clusters[name]
	if !ok {
		conf.Clusters[name] = &model.Cluster{URL: url}
		return StoreGlobalConfigToFile(conf, utils.GetGlobalConfigFile())
	}
	return errors.New(name + " already exists")
}

// GetClusterURL returns the url of the cluster
func GetClusterURL(name string) (string, error) {
	// Sanity check
	if len(name) == 0 {
		return "", errors.New("Cluster name needs to be provided")
	}

	// Load config from file
	conf, err := LoadGlobalConfigFromFile(utils.GetGlobalConfigFile())
	if err != nil {
		return "", err
	}

	c, ok := conf.Clusters[name]
	if !ok {
		return "", errors.New(name + " does not exist")
	}

	return c.URL, nil
}

// RemoveCluster removes a cluster from the config
func RemoveCluster(name string) error {

	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster name needs to be provided")
	}

	// Load config from file
	conf, err := LoadGlobalConfigFromFile(utils.GetGlobalConfigFile())
	if err != nil {
		return err
	}

	_, ok := conf.Clusters[name]
	if !ok {
		return errors.New(name + " does not exist")
	}

	delete(conf.Clusters, name)
	return StoreGlobalConfigToFile(conf, utils.GetGlobalConfigFile())
}

// SetClusterURL sets a cluster url
func SetClusterURL(name, url string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster name needs to be provided")
	}
	if len(url) == 0 {
		return errors.New("Cluster url needs to be provided")
	}

	// Load config from file
	conf, err := LoadGlobalConfigFromFile(utils.GetGlobalConfigFile())
	if err != nil {
		return err
	}

	conf.Clusters[name] = &model.Cluster{URL: url}
	return StoreGlobalConfigToFile(conf, utils.GetGlobalConfigFile())
}

// SetClusterToken sets a cluster token
func SetClusterToken(name, token string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster name needs to be provided")
	}

	// Load config from file
	conf, err := LoadGlobalConfigFromFile(utils.GetGlobalConfigFile())
	if err != nil {
		return err
	}

	c, p := conf.Clusters[name]
	if !p {
		return errors.New("Cluster not present")
	}

	c.Token = token
	return StoreGlobalConfigToFile(conf, utils.GetGlobalConfigFile())
}
