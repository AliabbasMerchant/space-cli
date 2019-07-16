package config

import (
	"fmt"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spaceuptech/space-cli/model"
	"github.com/spaceuptech/space-cli/utils"
)

// GenerateConfig runs the cli survey and generates a config file
func GenerateConfig() error {
	c := new(model.Deploy)

	// Ask project level details
	err := survey.AskOne(&survey.Input{Message: "name of app:"}, &c.Name, survey.Required)
	if err != nil {
		return err
	}
	err = survey.AskOne(&survey.Input{Message: "project name:"}, &c.Project, survey.Required)
	if err != nil {
		return err
	}
	err = survey.AskOne(&survey.Select{
		Message: "kind of app:",
		Options: []string{utils.KindService, utils.KindWebApp},
	}, &c.Kind, survey.Required)
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

	switch c.Kind {
	case utils.KindService:
		// Ask Runtime details
		c.Runtime = new(model.Runtime)
		name := ""
		err = survey.AskOne(&survey.Select{
			Message: "runtime type:",
			Options: []string{utils.Python3, utils.Java11, utils.Golang, utils.NodeJS},
		}, &name, survey.Required)
		if err != nil {
			return err
		}
		switch name {
		case utils.Python3:
			c.Runtime.Name = utils.Python3Image
		case utils.Java11:
			c.Runtime.Name = utils.Java11Image
		case utils.Golang:
			c.Runtime.Name = utils.GolangImage
		case utils.NodeJS:
			c.Runtime.Name = utils.NodeJSImage
		}

		switch c.Runtime.Name {
		case utils.Python3Image:
			err = survey.AskOne(&survey.Input{Message: "command to install:", Default: "pip install requirements.txt"}, &c.Runtime.Install, survey.Required)
			if err != nil {
				return err
			}
			err = survey.AskOne(&survey.Input{Message: "command to run:", Default: "python main.py"}, &c.Runtime.Run, survey.Required)
			if err != nil {
				return err
			}
		case utils.NodeJSImage:
			err = survey.AskOne(&survey.Input{Message: "command to install:", Default: "npm install"}, &c.Runtime.Install, survey.Required)
			if err != nil {
				return err
			}
			err = survey.AskOne(&survey.Input{Message: "command to run:", Default: "npm start"}, &c.Runtime.Run, survey.Required)
			if err != nil {
				return err
			}
		case utils.Java11Image:
			c.Runtime.Install = ""
			err = survey.AskOne(&survey.Input{Message: "command to run (please make a jar file):", Default: "java jarfile.jar"}, &c.Runtime.Run, survey.Required)
			if err != nil {
				return err
			}
		case utils.GolangImage:
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
		constraints := model.Constraints{}
		c.Constraints = &constraints
		err = survey.AskOne(&survey.Input{Message: "replicas", Default: "1"}, &constraints.Replicas, survey.Required)
		if err != nil {
			return err
		}
		var cpu float32 = 0.2
		err = survey.AskOne(&survey.Input{Message: "cpu limit", Default: "0.2"}, &cpu, survey.Required)
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
		// c.Clusters = make(map[string]string)

	case utils.KindWebApp:
		c.Runtime = new(model.Runtime)
		c.Runtime.Name = utils.WebAppImage
		c.Runtime.Install = ""
		c.Runtime.Run = ""

		c.Constraints = new(model.Constraints)
		c.Constraints.Replicas = 1
		cpu := float32(0.2)
		memory := int64(200)
		c.Constraints.CPU = &cpu
		c.Constraints.Memory = &memory

		fmt.Println("Enter port expose details")
		c.Expose = []*model.Expose{}
		err := addExpose(c)
		if err != nil {
			return err
		}
		name := "web"
		var port int32 = 8000
		c.Ports = append(c.Ports, &model.Port{Name: &name, Port: port})
		c.Env = map[string]string{"PREFIX": c.Expose[0].Prefix}
	}

	return StoreConfigToFile(c, utils.DefaultConfigFilePath)
}
