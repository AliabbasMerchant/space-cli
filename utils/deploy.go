package utils

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/spaceuptech/space-cli/model"
)

// Deploy deploys the archive and config to the space cloud clusters
func Deploy(name string, cluster *model.Cluster, conf *model.Deploy) error {
	// Check if user is logged in
	if len(cluster.Token) == 0 {
		return errors.New("Not logged in the cluster")
	}

	// Create an ignore object
	ignore, err := NewIgnore(conf.Ignore)
	if err != nil {
		log.Println(err)
		return err
	}

	// Get the file list to be archived
	files, err := ignore.GetFileList(conf.WorkingDir)
	if err != nil {
		return err
	}

	// Archive the files
	if err := ZipFiles(ZipName, files); err != nil {
		return err
	}
	defer os.Remove(ZipName)

	// Marshal the config to json
	json, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	if err := SendToCluster(cluster.Token, cluster.URL+"/v1/api/deploy", ZipName, json); err != nil {
		return err
	}

	return nil
}
