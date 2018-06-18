package dbs

import (
	"log"

	"github.com/EverLoSa/dbaas/dao"
	"github.com/EverLoSa/dbaas/docker"
	"github.com/EverLoSa/dbaas/model"
)

// CreateMySQLInstance creates a new instance of MySQL
func CreateMySQLInstance(db *model.DataBaseInfo) error {

	var err error
	// Creates a new container for MySql
	log.Println("Creating a docker container with image: mysql")
	if err = docker.NewContainer(db, "mysql"); err != nil {
		log.Println("An error ocurred creating the container")
		return err
	}

	// save in db.
	log.Println("Saving database instance in mongo...")
	if err := dao.MongoCreate(db); err != nil {
		log.Println("An error ocurred creating the DB in Mongo")
		return err
	}

	return err
}
