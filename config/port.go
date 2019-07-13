package config

import (
	"errors"
	"strconv"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spaceuptech/space-cli/utils"
	"github.com/spaceuptech/space-cli/model"
)

func askPortInfo(port *model.Port, defPort, defProtocol string) error {
	var p string
	err := survey.AskOne(&survey.Input{Message: "port:", Default: defPort}, &p, survey.Required)
	if err != nil {
		return err
	}
	pp, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return errors.New("Port must be an integer value")
	}
	port.Port = int32(pp)
	err = survey.AskOne(&survey.Input{Message: "protocol:", Default: defProtocol}, &p, nil)
	if err != nil {
		return err
	}
	if p != "" {
		port.Protocol = &p
	}

	return nil
}

// AddPort adds new port information
func AddPort(name string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Port name cannot be empty")
	}
	
	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	exists := false
	for _, p := range conf.Ports {
		if p.Name != nil {
			if *p.Name == name {
				exists = true
				break
			}
		}
	}
	if exists {
		return errors.New("Port " + name + " already exists")
	}
	port := &model.Port{Name: &name}
	err = askPortInfo(port, "8000", "TCP")
	if err != nil {
		return err
	}
	conf.Ports = append(conf.Ports, port)
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}

// RemovePort removes existing port information
func RemovePort(name string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Port name cannot be empty")
	}
	
	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	index := -1
	for i, p := range conf.Ports {
		if p.Name != nil {
			if *p.Name == name {
				index = i
				break
			}
		}
	}
	if index != -1 {
		conf.Ports = append(conf.Ports[:index], conf.Ports[index+1:]...)
		return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
	}
	return nil
}

// EditPort edits existing port information
func EditPort(name string) error {
	// Sanity check
	if len(name) == 0 {
		return errors.New("Port name cannot be empty")
	}
	
	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	exists := false
	for _, p := range conf.Ports {
		if p.Name != nil {
			if *p.Name == name {
				exists = true
				if p.Protocol != nil {
					askPortInfo(p, strconv.Itoa(int(p.Port)), *p.Protocol)
					break
				}
				askPortInfo(p, strconv.Itoa(int(p.Port)), "")
				break
			}
		}
	}
	if !exists {
		return errors.New("Port " + name + " does not exist")
	}
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}
