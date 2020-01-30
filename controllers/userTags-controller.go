package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Midlu/GoShit/Mine/TaskPad/db"
	"github.com/gorilla/mux"
)

type userTagsResponse struct {
	Status int          `json:"status"`
	Data   []db.UserTag `json:"data"`
	Error  string       `json:"error"`
}

//CreateUserTagRoutes make the routes for anything related to the tasks
func CreateUserTagRoutes(r *mux.Router, database *sql.DB) {
	//route for this one is: /api/tasks?pending=boolean
	r.HandleFunc("/api/usertags", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.GetUserTags(database)
		//there was an error trying to retrieve the records, stop here
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", err)
			return
		}
		//make the json response
		response, err := json.Marshal(userTagsResponse{Status: http.StatusOK, Data: rows})
		//if there's an error, stop here
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", err)
			return
		}
		//update headers and return this
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		log.Println("Finished sending data back.")
	}).Methods("GET")
}
