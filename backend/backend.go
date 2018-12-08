package backend

import (
	"encoding/json"
	"github.com/globalsign/mgo"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)
import "github.com/dsr373/friendex/myutil"

// TODO: replace manual error checks with myutil

// ConfigPath is the directory where the program's configuration is stored, usually $HOME/.config/friendex
var ConfigPath = myutil.ConfigDir()

// CredentialsFilename is the file where the database credentials are stored
var CredentialsFilename = filepath.Join(ConfigPath, "credentials.json")

// Credentials is a struct that stores all you need to connect to the database
// For MongoDB it's just a URI string
type Credentials struct {
	URI string
}

func loadCredentials(filename string) Credentials {
	log.Println("Reading database credentials...")
	var creds Credentials

	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	if err != nil {
		// if there is some error, log and exit
		log.Fatalf("Opening credentials file failed: %v", err)
	}
	// the credentials are stored locally and can be read
	log.Printf("Successfully Opened %s\n", filename)
	defer jsonFile.Close()

	// read everything and decode
	readBytes, readErr := ioutil.ReadAll(jsonFile)
	if readErr != nil {
		log.Fatalf("Reading credentials file failed: %v", readErr)
	}
	json.Unmarshal(readBytes, &creds)

	return creds
}

// OpenConnection function opens the connection to the database and returns the session object
func OpenConnection() *mgo.Session {
	creds := loadCredentials(CredentialsFilename)

	log.Println("Opening connection...")
	session, err := mgo.Dial(creds.URI)
	myutil.CheckErr(err, "Error creating session")
	databaseNames, _ := session.DatabaseNames()
	log.Printf("Connection opened successfully. Databases: %s\n", databaseNames)
	return session
}
