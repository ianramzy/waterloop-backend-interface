// using mgo package - https://github.com/globalsign/mgo
// Tests - https://github.com/globalsign/mgo/blob/master/session_test.go
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

//func connect()(*mgo.Session, ){
//
//	session, err := mgo.Dial(url)
//	if err != nil {
//		fmt.Printf("couldnt connect to DB :(\n")
//		log.Fatal(err)
//	}
//	defer session.Close()
//	fmt.Printf("Successfully connected to mongodb server at %v\n", url)
//
//	dbName := "test"
//	db := session.DB(dbName)
//	if db == nil {
//		fmt.Printf("db '%v' not found, exiting...\n", dbName)
//		return
//	}
//
//	// define locations of collections, just two for now: sensor and command
//	collectionSensor := session.DB("test").C("sensor")
//	collectionCommand := session.DB("test").C("command")
//
//}
func main() {
	session, err := mgo.Dial(url)

	if err != nil {
		fmt.Printf("couldnt connect to DB :(\n")
		log.Fatal(err)
	}
	defer session.Close()
	fmt.Printf("Successfully connected to mongodb server at %v\n", url)

	dbName := "test"
	db := session.DB(dbName)
	if db == nil {
		fmt.Printf("db '%v' not found, exiting...\n", dbName)
		return
	}

	// define locations of collections, just two for now: sensor and command
	collectionSensor := session.DB("test").C("sensor")
	collectionCommand := session.DB("test").C("command")

	testSD := SensorData{SensorID: 001, Data: 99, Timestamp: 123456789}
	testCD := CommandData{CommandID: 005, Data: 1234, Timestamp: 123456789}

	insertSensorItem(testSD, collectionSensor)
	insertCommandItem(testCD, collectionCommand)

	fmt.Printf("All done!\n")
}

func insertSensorItem(newItem SensorData, collection *mgo.Collection) {
	err := collection.Insert(&SensorData{SensorID: newItem.SensorID, Data: newItem.Data, Timestamp: newItem.Timestamp})
	if err != nil {
		fmt.Printf("Error inserting sensor item\n")
	} else {
		fmt.Printf("Sensor item Inserted\n")
	}
}

func insertCommandItem(newItem CommandData, collection *mgo.Collection) {
	err := collection.Insert(&SensorData{SensorID: newItem.CommandID, Data: newItem.Data, Timestamp: newItem.Timestamp})
	if err != nil {
		fmt.Printf("Error inserting command item\n")
	} else {
		fmt.Printf("command item Inserted\n")
	}
}

//func listDocs(db *mgo.Database, col string) {
//
//	coll := db.C(col)
//
//	if coll == nil {
//		return
//	}
//
//	type Document struct {
//		ID   int    `json:"_id,omitempty"`
//		Desc string `json:"desc,omitempty"`
//		Done bool   `json:"done,omitempty"`
//	}
//
//	var result []map[string]string // see bson.M
//
//	coll.Find(nil).All(&result)
//
//	for i, d := range result {
//		fmt.Printf("\tDoc%3v - %#v\n", i+1, d)
//		obj := bson.ObjectId(d["_id"])
//		fmt.Printf("\t\tHex: %v, String: %v, Time: %v\n", obj.Hex(), obj.String(), obj.Time())
//	}
//}

//func main() {
//	dbName := "test"
//	if 1 == len(os.Args) {
//		log.Warnf("No db specified, using '%v'", dbName)
//	} else {
//		dbName = os.Args[1]
//	}
//	// list documents in selected database collections
//	// ----
//	session, err := mgo.Dial(url)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer session.Close()
//
//	log.Infof("Successfully connected to mongodb server at %v", url)
//
//	db := session.DB(dbName)
//	if db == nil {
//		log.Errorf("db '%v' not found, exiting...", dbName)
//		return
//	}
//
//	// iterate collections
//	fmt.Printf("Collections in db '%v':\n", dbName)
//	cols, err := db.CollectionNames()
//	if err != nil {
//		return
//	}
//
//	for _, c := range cols {
//		fmt.Printf("[%v]\n", c)
//		listDocs(db, c)
//	}
//}
//
//func listDocs(db *mgo.Database, col string) {
//	coll := db.C(col)
//	if coll == nil {
//		return
//	}
//
//	type Document struct {
//		ID   int    `json:"_id,omitempty"`
//		Desc string `json:"desc,omitempty"`
//		Done bool   `json:"done,omitempty"`
//	}
//
//	var result []map[string]string // see bson.M
//	coll.Find(nil).All(&result)
//
//	for i, d := range result {
//		fmt.Printf("\tDoc%3v - %#v\n", i+1, d)
//		obj := bson.ObjectId(d["_id"])
//		fmt.Printf("\t\tHex: %v, String: %v, Time: %v\n", obj.Hex(), obj.String(), obj.Time())
//	}
//}
