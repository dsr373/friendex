package backend

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/dsr373/friendex/myutil"
)

// DefaultCtx is the default context for executing queries, it times out after 10 seconds
var defaultCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)

// ConfigPath is the directory where the program's configuration is stored, usually $HOME/.config/friendex
var ConfigPath = myutil.ConfigDir()

// CredentialsFilename is the file where the database credentials are stored
var CredentialsFilename = filepath.Join(ConfigPath, "credentials.json")

// Credentials is a struct that stores all you need to connect to the database
// For MongoDB it's just a URI string
type Credentials struct {
	URI string
}

// Transaction defines transactions duh
type Transaction struct {
	ID          int
	PayeeID     int
	ReceiverIds []int
	Amount      float64
}

// User defines users duh
type User struct {
	ID      int
	Name    string
	Balance float64
}

func loadCredentials(filename string) Credentials {
	log.Println("Reading database credentials...")
	var creds Credentials

	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	myutil.CheckErr(err, "Opening credentials file failed")

	// the credentials are stored locally and can be read
	log.Printf("Successfully Opened %s\n", filename)
	defer jsonFile.Close()

	// read everything and decode
	readBytes, readErr := ioutil.ReadAll(jsonFile)
	myutil.CheckErr(readErr, "Reading credentials file failed")
	json.Unmarshal(readBytes, &creds)

	return creds
}

// OpenClient provides a new client connection to the database
func OpenClient() *mongo.Client {
	creds := loadCredentials(CredentialsFilename)

	log.Printf("Opening connection...")
	client, err := mongo.Connect(defaultCtx, creds.URI)
	myutil.CheckErr(err, "Error creating client")
	log.Printf("Connection opened successfully.")
	return client
}

// InsertUser adds a new user to the database
func InsertUser(cl *mongo.Client, user User) error {
	log.Printf("Inserting new user: %v", user)

	coll := cl.Database("friendex").Collection("users")
	_, err := coll.InsertOne(defaultCtx, user)
	return err
}

// InsertTransaction puts a new transaction into the database
func InsertTransaction(cl *mongo.Client, tr Transaction) error {
	log.Printf("Inserting new transaction: %v", tr)

	// TODO: Check the transaction for integrity

	coll := cl.Database("friendex").Collection("transactions")
	_, err := coll.InsertOne(defaultCtx, tr)
	return err
}
