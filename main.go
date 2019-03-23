// using mgo package - https://github.com/globalsign/mgo
package main

import (
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

const (
	url = "localhost"
)

type SensorData struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	SensorID  uint32
	Data      uint64
	Timestamp uint64
}

type CommandData struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	CommandID uint32
	Data      uint64
	Timestamp uint64
}

func main() {
	session, collectionSensor, collectionCommand := connect()
	defer session.Close()

	// insert test data, this is just for testing purposes
	testSD := SensorData{SensorID: 001, Data: 99, Timestamp: 123456789}
	insertSensorItem(testSD, collectionSensor)
	testCD := CommandData{CommandID: 005, Data: 1234, Timestamp: 123456789}
	insertCommandItem(testCD, collectionCommand)
}

// function used to connect to MongoDB server
func connect() (*mgo.Session, *mgo.Collection, *mgo.Collection) {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("couldnt connect to DB :(\n")
		log.Fatal(err)
	}
	fmt.Printf("Successfully connected to mongodb server at %v\n", url)
	dbName := "test" //change to your DB name
	db := session.DB(dbName)
	if db == nil {
		fmt.Printf("db '%v' not found, exiting...\n", dbName)
		log.Fatal()
	}
	// define locations of collections, just two for now: sensor and command
	collectionSensor := session.DB("test").C("sensor")
	collectionCommand := session.DB("test").C("command")
	return session, collectionSensor, collectionCommand
}

func insertSensorItem(newItem SensorData, collection *mgo.Collection) {
	err := collection.Insert(&SensorData{SensorID: newItem.SensorID, Data: newItem.Data, Timestamp: newItem.Timestamp})
	if err != nil {
		fmt.Printf("Error inserting sensor item\n")
	} else {
		fmt.Printf("Sensor Inserted\n")
	}
}

func insertCommandItem(newItem CommandData, collection *mgo.Collection) {
	err := collection.Insert(&SensorData{SensorID: newItem.CommandID, Data: newItem.Data, Timestamp: newItem.Timestamp})
	if err != nil {
		fmt.Printf("Error inserting command item\n")
	} else {
		fmt.Printf("Command Inserted\n")
	}
}
