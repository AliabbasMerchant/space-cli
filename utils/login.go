package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spaceuptech/space-cli/model"
)

// GenerateCredentials run the cli survey to generate user credentials
func GenerateCredentials() (*model.Credentials, error) {
	c := new(model.Credentials)

	err := survey.AskOne(&survey.Input{Message: "Enter Username:"}, &c.User, survey.Required)
	if err != nil {
		return nil, err
	}
	err = survey.AskOne(&survey.Password{Message: "Enter Password:"}, &c.Pass, survey.Required)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Login logs the user in
func Login(cluster, user, pass string) error {
	if cluster == "" {
		return errors.New("Cluster url needs to be provided")
	}

	var c *model.Credentials
	if user == "" || pass == "" {
		cTemp, err := GenerateCredentials()
		if err != nil {
			return err
		}
		c = cTemp
	} else {
		c = &model.Credentials{user, pass}
	}

	return loginRequest(cluster, c)
}

type loginResponse struct {
	Token string `json:"token"`
}

func loginRequest(cluster string, c *model.Credentials) error {
	url := cluster + "/v1/api/config/login"

	data, _ := json.Marshal(c)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	obj := new(loginResponse)
	if err := json.NewDecoder(res.Body).Decode(obj); err != nil {
		return err
	}

	return writeToken(obj.Token)
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
func writeToken(token string) error {
	homeDir := userHomeDir()
	dir := homeDir + "/.space-cloud/cli"
	os.MkdirAll(dir, os.ModePerm)
	return ioutil.WriteFile(dir+"/token.txt", []byte(token), 0644)
}

func readToken() (string, error) {
	homeDir := userHomeDir()
	file := homeDir + "/.space-cloud/cli/token.txt"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
