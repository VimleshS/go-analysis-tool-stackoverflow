package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	APP_ID         = "APP_ID"
	API_KEY_ID     = "API_KEY_ID"
	API_KEY_SECRET = "API_KEY_SECRET"
	REDIRECT_URI   = "REDIRECT_URI"
)

var Config map[string]string

func init() {
	Config = make(map[string]string)
	ReadCredentialFromFile("app.properties")
}

func ReadCredentialFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, "=")
		Config[strings.TrimSpace(data[0])] = strings.TrimSpace(data[1])
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
