package utils

import (
	"os"
	"runtime"
)

func GetGlobalConfigFile() string {
	homeDir := userHomeDir()
	return homeDir + "/.space-cloud/cli/config.json"
}

func GetGlobalConfigDir() string {
	homeDir := userHomeDir()
	return homeDir + "/.space-cloud/cli"
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
