package model

import "github.com/globalsign/mgo/bson"

// DBSize size for provisioning.
type DBSize struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Big    string `json:"Big"`
}

// DataBaseInfo model that is stored in Mongo.
type DataBaseInfo struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	Type          string        `bson:"type" json:"type"`
	Size          string        `bson:"size" json:"size"`
	ContainerInfo struct {
		ContainerID string `bson:"containerid" json:"containerid"`
		Image       string `bson:"image" json:"image"`
		Port        string `bson:"port" json:"port"`
	}
}
