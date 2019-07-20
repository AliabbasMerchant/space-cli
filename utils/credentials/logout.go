package credentials

import (
	"errors"

	"github.com/spaceuptech/space-cli/config"
	"github.com/spaceuptech/space-cli/utils"
)

// Logout logs the user out of a given cluster
func Logout(name string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Cluster name needs to be provided")
	}

	// Load the global config
	c, err := config.LoadGlobalConfigFromFile(utils.GetGlobalConfigFile())
	if err != nil {
		return err
	}

	cluster, p := c.Clusters[name]
	if !p {
		return errors.New("Cluster not present")
	}
	cluster.Token = ""

	return config.StoreGlobalConfigToFile(c, utils.GetGlobalConfigFile())
}
