package config

import (
	"strconv"
	"errors"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spaceuptech/space-cli/utils"
	"github.com/spaceuptech/space-cli/model"
)

// AddExpose adds an exposed port and its proxy information
func AddExpose() error {
	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}
	
	err = addExpose(conf)
	if err != nil {
		return err
	}

	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}

func addExpose(conf *model.Deploy) error {
	expose := new(model.Expose)
	prefix := ""
	err := survey.AskOne(&survey.Input{Message: "prefix:"}, &prefix, survey.Required)
	if err != nil {
		return err
	}
	expose.Prefix = &prefix
	host := ""
	err = survey.AskOne(&survey.Input{Message: "host:"}, &host, survey.Required)
	if err != nil {
		return err
	}
	expose.Host = &host
	proxy := ""
	err = survey.AskOne(&survey.Input{Message: "proxy:"}, &proxy, survey.Required)
	if err != nil {
		return err
	}
	expose.Proxy = &proxy
	port := ""
	err = survey.AskOne(&survey.Input{Message: "port:"}, &port, survey.Required)
	if err != nil {
		return err
	}
	p, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		return errors.New("port must be an integer value")
	}
	expose.Port = int32(p)
	conf.Expose = append(conf.Expose, expose)

	return nil
}