package config

import (
	"errors"
	"strconv"

	"github.com/spaceuptech/space-cli/utils"
)

// SetReplicas sets the number of replicas
func SetReplicas(replicas string) error {
	// Sanity check
	if len(replicas) == 0 {
		return errors.New("Number of replicas cannot be empty")
	}
	
	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	rep, err := strconv.ParseInt(replicas, 10, 32)
	if err != nil {
		return errors.New("Number of replicas must be an integer value")
	}
	if rep <= 0 {
		return errors.New("Number of replicas must be a positive integer")
	}
	
	conf.Constraints.Replicas = int32(rep)
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}

// SetCPU sets the cpu limit
func SetCPU(cpu string) error {
	// Sanity check
	if len(cpu) == 0 {
		return errors.New("CPU limit cannot be empty")
	}
	
	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	CPU, err := strconv.ParseFloat(cpu, 32)
	if err != nil {
		return errors.New("CPU limit must be a decimal value")
	}
	if CPU <= 0 {
		return errors.New("CPU limit must be a positive decimal")
	}
	c := float32(CPU)
	conf.Constraints.CPU = &c
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}

// SetMemory sets the memory limit
func SetMemory(memory string) error {
	// Sanity check
	if len(memory) == 0 {
		return errors.New("Memory limit cannot be empty")
	}
	
	// Load config from file
	conf, err := LoadConfigFromFile(utils.DefaultConfigFilePath)
	if err != nil {
		return err
	}

	mem, err := strconv.ParseInt(memory, 10, 64)
	if err != nil {
		return errors.New("Memory limit must be an integer value")
	}
	if mem <= 0 {
		return errors.New("Memory limit must be a positive integer")
	}
	
	conf.Constraints.Memory = &mem
	return StoreConfigToFile(conf, utils.DefaultConfigFilePath)
}