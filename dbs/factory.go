package dbs

import (
	"errors"
	"log"
	"strings"

	"github.com/EverLoSa/dbaas/dao"
	"github.com/EverLoSa/dbaas/docker"
	"github.com/EverLoSa/dbaas/model"
	"github.com/globalsign/mgo/bson"
)

// NewInstace creates a new dbaas instace of the specified type
func NewInstace(databse *model.DataBaseInfo) error {

	databse.ID = bson.NewObjectId()
	// Check de type to create the instace.
	switch strings.ToLower(databse.Type) {

	case "mysql":
		log.Println("A new instance of mysql will be created.")
		err := CreateMySQLInstance(databse)
		return err

	case "mariadb":
		log.Println("A new instance of mariadb will be created.")
		err := CreateMariaDBInstance(databse)
		return err
	default:
		err := errors.New("Unsupported type; Current supported types: mysql, mariadb")
		return err
	}
}

// DeleteInstance deletes a instance onf db in dbaas.
func DeleteInstance(id string) error {

	var err error
	var dataBaseInfo model.DataBaseInfo

	// Get instance info to delete.
	if err = dao.MongoGetDBByID(id, &dataBaseInfo); err != nil {
		log.Println("An error ocurred getting instance's info")
		return err
	}

	// Detelete instance from docker
	if err := docker.DeleteContainer(dataBaseInfo.ContainerInfo.ContainerID); err != nil {
		log.Println("An error ocurred deleting the Container")
		return err
	}

	// Delete instance from mongo
	if err := dao.MongoDeleteDB(id); err != nil {
		log.Println("An error ocurred deleted the instance from MongoDB")
		return err
	}
	return err

}
