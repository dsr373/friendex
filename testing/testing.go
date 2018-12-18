package testing

import (
	"github.com/dsr373/friendex/backend"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var jDoe = backend.User{ID: 1, Name: "John Doe", Balance: -3.50}
var aMic = backend.User{ID: 2, Name: "Archangel Michael", Balance: 1000}
var lTit = backend.User{ID: 3, Name: "Titus Livius", Balance: 42}

// PutFakeUsers inserts the fake users one by one into the database
func PutFakeUsers(cl *mongo.Client) error {
	err := backend.InsertUser(cl, jDoe)
	if err != nil {
		return err
	}
	err = backend.InsertUser(cl, aMic)
	if err != nil {
		return err
	}
	err = backend.InsertUser(cl, lTit)
	return err
}
