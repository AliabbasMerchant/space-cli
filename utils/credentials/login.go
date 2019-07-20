package credentials

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spaceuptech/space-cli/config"
	"github.com/spaceuptech/space-cli/model"
)

// GenerateCredentials run the cli survey to generate user credentials
func GenerateCredentials() (*model.Credentials, error) {
	c := new(model.Credentials)

	err := survey.AskOne(&survey.Input{Message: "username:"}, &c.User, survey.Required)
	if err != nil {
		return nil, err
	}
	err = survey.AskOne(&survey.Password{Message: "password:"}, &c.Pass, survey.Required)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Login logs the user in
func Login(cluster, user, pass string) error {
	if cluster == "" {
		return errors.New("Cluster name needs to be provided")
	}

	url, err := config.GetClusterURL(cluster)
	if err != nil {
		return err
	}

	var c *model.Credentials
	if user == "" || pass == "" {
		cTemp, err := GenerateCredentials()
		if err != nil {
			return err
		}
		c = cTemp
	} else {
		c = &model.Credentials{User: user, Pass: pass}
	}

	token, err := loginRequest(url, c)
	if err != nil {
		return err
	}

	return config.SetClusterToken(cluster, token)
}

type loginResponse struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

func loginRequest(cluster string, c *model.Credentials) (string, error) {
	url := cluster + "/v1/api/config/login"

	data, _ := json.Marshal(c)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	obj := new(loginResponse)
	if err := json.NewDecoder(res.Body).Decode(obj); err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", errors.New(obj.Error)
	}

	return obj.Token, nil
}
