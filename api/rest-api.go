package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/EverLoSa/dbaas/dao"
	"github.com/EverLoSa/dbaas/dbs"
	"github.com/EverLoSa/dbaas/model"
	"github.com/gorilla/mux"
)

// CreateDataBaseEndpoint creates a new data base.
func CreateDataBaseEndpoint(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	var database model.DataBaseInfo

	log.Println("Reading request body...")
	if err := json.NewDecoder(r.Body).Decode(&database); err != nil {
		// send error response because the request could not be parsed.
		SendResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	log.Println("Creating a new instance of database in dbaas...")
	if err := dbs.NewInstace(&database); err != nil {
		SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	fmt.Fprintln(w, "Database has been created successfully!")
	SendResponse(w, http.StatusOK, database)
}

// ListDataBasesEndpoint lists all databases.
func ListDataBasesEndpoint(w http.ResponseWriter, r *http.Request) {

	log.Println("Getting Databases...")
	databases, err := dao.MongoListAll()
	if err != nil {
		// send error response because the list of dbs could not be retreived from mongo.
		SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	SendResponse(w, http.StatusOK, databases)
}

// DeleteDatabaseEndpoint deletes a database.
func DeleteDatabaseEndpoint(w http.ResponseWriter, r *http.Request) {

	log.Println("Reading request body...")
	params := mux.Vars(r)

	log.Println("Deleting instance from dbaas")
	log.Println(params["id"])
	if err := dbs.DeleteInstance(params["id"]); err != nil {
		SendResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	SendResponse(w, http.StatusOK, "Database has been successfully deleted!")

}

// SendResponse sends feedback to the client with the request status
func SendResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {

	log.Println("Starting DbaaS REST API....")
	router := mux.NewRouter()
	router.HandleFunc("/dbaas/", ListDataBasesEndpoint).Methods("GET")
	router.HandleFunc("/dbaas/", CreateDataBaseEndpoint).Methods("POST")
	router.HandleFunc("/dbaas/{id}", DeleteDatabaseEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8888", router))

}
