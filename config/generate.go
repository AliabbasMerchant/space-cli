package config

import (
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"

	"gitlab.com/spaceuptech/space-registry/utils"
)

// GenerateConfig runs the cli survey and generates a config file
func GenerateConfig() error {
	c := new(Deploy)

	// Ask project leve details
	err := survey.AskOne(&survey.Input{Message: "name:"}, &c.Name, survey.Required)
	if err != nil {
		return err
	}
	err = survey.AskOne(&survey.Input{Message: "project name:"}, &c.Project, survey.Required)
	if err != nil {
		return err
	}

	// Ask working environment details
	err = survey.AskOne(&survey.Input{Message: "working directory:", Default: "."}, &c.WorkingDir, survey.Required)
	if err != nil {
		return err
	}
	err = survey.AskOne(&survey.Input{Message: "ignore file:", Default: ".gitignore"}, &c.Ignore, survey.Required)
	if err != nil {
		return err
	}
	c.Ignore = strings.TrimSpace(c.Ignore)

	// Ask Runtime details
	c.Runtime = new(Runtime)
	err = survey.AskOne(&survey.Select{
		Message: "runtime type:",
		Options: []string{utils.Python3, utils.Java11, utils.Golang, utils.NodeJS},
	}, &c.Runtime.Name, survey.Required)
	if err != nil {
		return err
	}

	switch c.Runtime.Name {
	case utils.Python3:
		err = survey.AskOne(&survey.Input{Message: "command to install:", Default: "pip install requirements.txt"}, &c.Runtime.Install, survey.Required)
		if err != nil {
			return err
		}
		err = survey.AskOne(&survey.Input{Message: "command to run:", Default: "python main.py"}, &c.Runtime.Run, survey.Required)
		if err != nil {
			return err
		}
	case utils.NodeJS:
		err = survey.AskOne(&survey.Input{Message: "command to install:", Default: "npm install"}, &c.Runtime.Install, survey.Required)
		if err != nil {
			return err
		}
		err = survey.AskOne(&survey.Input{Message: "command to run:", Default: "npm start"}, &c.Runtime.Run, survey.Required)
		if err != nil {
			return err
		}
	case utils.Java11:
		c.Runtime.Install = ""
		err = survey.AskOne(&survey.Input{Message: "command to run (please make a jar file):", Default: "java jarfile.jar"}, &c.Runtime.Run, survey.Required)
		if err != nil {
			return err
		}
	case utils.Golang:
		err = survey.AskOne(&survey.Input{Message: "command to setup:", Default: "go build"}, &c.Runtime.Install, survey.Required)
		if err != nil {
			return err
		}
		err = survey.AskOne(&survey.Input{Message: "command to run:", Default: "./executable"}, &c.Runtime.Run, survey.Required)
		if err != nil {
			return err
		}
	}

	// Ask contraint details
	constraints := Constraints{}
	c.Constraints = &constraints
	err = survey.AskOne(&survey.Input{Message: "replicas", Default: "1"}, &constraints.Replicas, survey.Required)
	if err != nil {
		return err
	}
	var cpu float32 = 0.1
	err = survey.AskOne(&survey.Input{Message: "cpu limit", Default: "0.1"}, &cpu, survey.Required)
	if err != nil {
		return err
	}
	constraints.CPU = &cpu

	var memory int64 = 200
	err = survey.AskOne(&survey.Input{Message: "memory limit (in MBs)", Default: "200"}, &memory, survey.Required)
	if err != nil {
		return err
	}
	constraints.Memory = &memory
	c.Env = map[string]string{"exampleEnvVariable": "exampleValue"}
	// c.Clusters = make(map[string]string)

	return StoreConfigToFile(c, utils.DefaultConfigFilePath)
}
