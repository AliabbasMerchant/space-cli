package utils

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func newfileUploadRequest(uri, zipPath string, config []byte) (*http.Request, error) {
	r, w := io.Pipe()
	m := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		defer m.Close()

		zipFile, err := os.Open(zipPath)
		if err != nil {
			log.Println(err)
		}
		defer zipFile.Close()
		// configFile, err := os.Open(configPath)
		// if err != nil {
		// 	log.Println(err)
		// }
		// defer configFile.Close()

		zipPart, err := m.CreateFormFile("zip", filepath.Base(zipPath))
		if err != nil {
			log.Println(err)
		}
		_, err = io.Copy(zipPart, zipFile)
		if err != nil {
			log.Println(err)
		}
		configPart, err := m.CreateFormFile("config", DefaultConfigFilePath)
		if err != nil {
			log.Println(err)
		}
		// _, err = io.Copy(configPart, configFile)
		configPart.Write(config)
	}()

	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		return nil, err
	}

	// Read the token from the file
	token, err := readToken()
	if err != nil {
		return nil, errors.New("You are not logged in to the cluster")
	}

	// Set the appropriate headers
	req.Header.Set("Content-Type", m.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+token)
	return req, err
}

// SendToCluster sends the config & archive to the space cloud cluster
func SendToCluster(url, zip string, conf []byte) error {
	request, err := newfileUploadRequest(url, zip, conf)
	if err != nil {
		return err
	}

	// Make a new http client and fire the request
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Marshal the response
	obj := map[string]string{}
	if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
		return err
	}

	if res.StatusCode == http.StatusOK {
		return nil
	}

	return errors.New(obj["error"])

}
