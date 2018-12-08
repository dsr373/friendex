package main

import (
	"encoding/json"
	"fmt"
	"os"
)
import "github.com/dsr373/friendex/backend"
import "github.com/dsr373/friendex/myutil"

// checkCredentials asks for new database log in information if it is not already present
// it is in main rather than on the backend because it is part of the interface
func checkCredentials() {
	// try to open the file
	filename := backend.CredentialsFilename
	_, err := os.Open(filename)

	// if there is no file, then ask for credentials and store them in a file
	if os.IsNotExist(err) {
		fmt.Print("Cannot find database connection info. Please input URI: ")

		// read standard input and put it into creds
		var uri string
		fmt.Scanln(&uri)
		creds := backend.Credentials{URI: uri}
		credsMarshalled, _ := json.Marshal(creds)

		// open a new file and write stuff to it
		mkdirErr := os.MkdirAll(backend.ConfigPath, os.ModePerm)
		myutil.CheckErr(mkdirErr, "Error creating config directory")
		jsonFile, openErr := os.Create(filename)
		myutil.CheckErr(openErr, "Error creating new credentials file")
		defer jsonFile.Close()
		_, writeErr := jsonFile.Write(credsMarshalled)
		myutil.CheckErr(writeErr, "Error writing to credentials file")
	}
}

func main() {
	checkCredentials()
	session := backend.OpenConnection()
	defer session.Close()
}
