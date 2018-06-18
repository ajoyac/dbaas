package dbs

import (
	"testing"
	"log"

	"github.com/EverLoSa/dbaas/model"
)

func TestNewInstance(t *testing.T) {


	database := model.DataBaseInfo{Name:"Test1",Type:"mysql",Size:"small"}
	NewInstace(&database)

	log.Println(database)
	if database.Name == "" {
		t.Fatalf("Test TestNewInstance:Name failed")
		t.Log("tons")
	}

	if database.Type == "" {
		t.Fatalf("Test TestNewInstance:Type failed")
		t.Log()
	}

	if database.Size == "" {
		t.Fatalf("Test TestNewInstance:Size failed")
		t.Log()
	}


}
