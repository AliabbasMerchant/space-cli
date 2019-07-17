package config

import (
	"fmt"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spaceuptech/space-cli/model"
	"github.com/spaceuptech/space-cli/utils"
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
	err := survey.AskOne(&survey.Input{Message: "url prefix to match against:", Default: "/"}, &prefix, survey.Required)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	expose.Prefix = prefix
	host := ""
	err = survey.AskOne(&survey.Input{Message: "hostname to match against:", Default: ""}, &host, nil)
	if err != nil {
		return err
	}
	expose.Host = host

	expose.Proxy = fmt.Sprintf("http://%s:8000%s", conf.Name, prefix)
	conf.Expose = append(conf.Expose, expose)

	return nil
}
