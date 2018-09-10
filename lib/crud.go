package lib

import (
	"fmt"
	// package to connect with mongo db
	"gopkg.in/mgo.v2"
	// package bson
	"gopkg.in/mgo.v2/bson"
)

// model for data region
type Region struct {
	// bson property is to define name on field on collection db
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Name    string        `bson:"name,omitempty"`
	Number  string        `bson:"number"`
	EventID string        `bson:"event_id"`
}

// method to create session to connect on database
func connect() (*mgo.Session, error) {
	// create connection object session
	// parameter is connection string from server mongodb
	var session, err = mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}

	return session, nil
}

// find method
func Find() {
	// create session
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}
	// close session
	defer session.Close()

	var collection = session.DB("sibiti").C("regions")

	var region Region
	// var selector = bson.M{"name": "Indonesia"}
	err = collection.Find(bson.M{"name": "Indonesia"}).One(&region)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("Name : ", region.Name)
	fmt.Println("Number : ", region.Number)
	fmt.Println("Event Id : ", region.EventID)
	fmt.Println(region.ID)
}
