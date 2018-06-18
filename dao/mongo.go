package dao

import (
	"log"

	"github.com/EverLoSa/dbaas/model"
	"github.com/globalsign/mgo/bson"
	"gopkg.in/mgo.v2"
)

// mongo collection
const (
	Collection   = "dbaas_instances"
	MongoDBHosts = "127.0.0.1:27017"
	Database     = "dbaas"
)

// DatabasesDAO struct
type DatabasesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// MongoConnect : connects with mongo: databases collection
func MongoConnect() {

	dbs := DatabasesDAO{Server: MongoDBHosts, Database: Database}
	var err error
	session, err := mgo.Dial(dbs.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dbs.Database)

}

// MongoCreate conection with mongodb
func MongoCreate(database *model.DataBaseInfo) error {
	MongoConnect()
	err := db.C(Collection).Insert(&database)
	return err
}

// method to close connection.

// MongoListAll the databases in the dbaas collection.
func MongoListAll() ([]model.DataBaseInfo, error) {
	MongoConnect()
	var databases []model.DataBaseInfo
	err := db.C(Collection).Find(bson.M{}).All(&databases)
	return databases, err
}

// MongoDeleteDB deletes a database instance from Mongo.
func MongoDeleteDB(id string) error {
	MongoConnect()
	err := db.C(Collection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

// MongoGetDBByID gets information by id
func MongoGetDBByID(id string, dataBaseInfo *model.DataBaseInfo) error {
	MongoConnect()
	err := db.C(Collection).FindId(bson.ObjectIdHex(id)).One(dataBaseInfo)
	return err
}
