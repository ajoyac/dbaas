package docker

import (
	"testing"

	"github.com/EverLoSa/dbaas/model"
)

func TestNewContainerMySQL(t *testing.T){

	database := model.DataBaseInfo{Name:"Test1",Type:"mysql",Size:"small"}
	NewContainer(&database,"mysql")

	if database.ContainerInfo.ContainerID == "" {
		t.Fatalf("Test TestNewInstance:ContainerID failed")
	}

	if database.ContainerInfo.Image == "" {
		t.Fatalf("Test TestNewInstance:Image failed")
	}

	if database.ContainerInfo.Port == "" {
		t.Fatalf("Test TestNewInstance:Port failed")
	}
	DeleteContainer(database.ContainerInfo.ContainerID)

}

func TestNewContainerMariaDB(t *testing.T){

	database := model.DataBaseInfo{Name:"Test1",Type:"mysql",Size:"small"}
	NewContainer(&database,"mariadb")

	if database.ContainerInfo.ContainerID == "" {
		t.Fatalf("Test TestNewInstance:ContainerID failed")
	}

	if database.ContainerInfo.Image == "" {
		t.Fatalf("Test TestNewInstance:Image failed")
	}

	if database.ContainerInfo.Port == "" {
		t.Fatalf("Test TestNewInstance:Port failed")
	}

	DeleteContainer(database.ContainerInfo.ContainerID)

}
