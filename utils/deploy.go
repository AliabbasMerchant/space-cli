package utils

import (
	"encoding/json"
	"log"
	"os"

	"gitlab.com/spaceuptech/space-registry/config"
)

// Deploy deploys the archive and config to the space cloud clusters
func Deploy() error {
	// Load config from file
	conf, err := config.LoadConfigFromFile(DefaultConfigFilePath)
	if err != nil {
		return err
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

	// Marshal the config to json
	json, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	// Deploy to all clusters
	for name, url := range conf.Clusters {
		err = SendToCluster(url+"/v1/api/deploy", ZipName, json)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Successfully deployed to cluster: " + name)
		}
	}
	return os.Remove(ZipName)
}
