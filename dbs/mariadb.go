package dbs

import (
	"log"

	"github.com/EverLoSa/dbaas/dao"
	"github.com/EverLoSa/dbaas/docker"
	"github.com/EverLoSa/dbaas/model"
)

// CreateMariaDBInstance creates a new instance of mariadb
func CreateMariaDBInstance(db *model.DataBaseInfo) error {

	var err error
	// Creates a new container for mariadb
	log.Println("Creating a docker container with image: mariadb")
	if err = docker.NewContainer(db, "mariadb"); err != nil {
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
