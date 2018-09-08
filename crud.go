package main

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
	ID       bson.ObjectId `bson:"_id,omitempty"`
	name     string        `bson:"name"`
	number   string        `bson:"number"`
	event_id string        `bson:"event_id"`
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

func main() {
	find()
}

// find method
func find() {
	// create session
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Println("1 session created")

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

	fmt.Println("Name :", region.name)
	fmt.Println(region.event_id)
}
