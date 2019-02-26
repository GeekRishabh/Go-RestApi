package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//user represents a sample database entity.
type User struct {
	ID   string `json:"id" `
	Name string `json:"name"`
}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("localhost/api_db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("api_db")
}

func collection() *mgo.Collection {
	return db.C("users")
}

// GetAll returns all users from the database.
func GetAll() ([]User, error) {
	res := []User{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetOne returns a single user from the database.
func GetOne(id string) (*User, error) {
	res := User{}

	if err := collection().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Save inserts an user to the database.
func Save(user User) error {
	return collection().Insert(user)
}

// Remove deletes an user from the database
func Remove(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
