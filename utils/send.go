package utils

import (
	"bytes"
	"io"
	"log"
	"errors"
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
		configPart, err := m.CreateFormFile("config", DefaultConfigFilePath)
		if err != nil {
			log.Println(err)
		}
		// _, err = io.Copy(configPart, configFile)
		configPart.Write(config)
	}()

	req, err := http.NewRequest("POST", uri, r)
	req.Header.Set("Content-Type", m.FormDataContentType())
	return req, err
}

func SendToCluster(url, zip string, conf []byte) error {
	request, err := newfileUploadRequest(url, zip, conf)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	} else {
		log.Println(resp.StatusCode)
		// if resp.StatusCode == http.StatusOK {
		// 	return nil
		// }
		body := &bytes.Buffer{}
		log.Println(body)
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			return err
		}
		if resp.StatusCode == http.StatusOK {
			return nil
		}
		resp.Body.Close()
		return errors.New(body.String())
	}
}