package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Midlu/GoShit/Mine/TaskPad/db"
	"github.com/Midlu/GoShit/Mine/TaskPad/utils"
	"github.com/gorilla/mux"
)

type taskResponse struct {
	Status int       `json:"status"`
	Data   []db.Task `json:"data"`
	Error  string    `json:"error"`
}

//CreateTaskRoutes make the routes for anything related to the tasks
func CreateTaskRoutes(r *mux.Router, database *sql.DB) {
	r.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		vars := r.URL.Query()
		rows, err := db.GetPendingTasks(database, vars)
		//there was an error trying to retrieve the records, stop here
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", err)
			return
		}
		//return a response to the front
		utils.RespondWithJSON(w, http.StatusOK, taskResponse{Status: http.StatusOK, Data: rows})
	}).Methods("GET")

	r.HandleFunc("/api/tasks", func(w http.ResponseWriter, r *http.Request) {
		//define the task Struct
		var newTask db.Task
		//make a decorder to read the body
		decoder := json.NewDecoder(r.Body)
		//decode the body into the new task struct, if there's an error return
		if err := decoder.Decode(&newTask); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", err)
			return
		}
		//why do this???
		defer r.Body.Close()
		dbInsertErr := db.CreateTask(database, newTask)
		if dbInsertErr != nil {
			http.Error(w, dbInsertErr.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", dbInsertErr)
			return
		}
		//return a response to the front
		utils.RespondWithJSON(w, http.StatusCreated, taskResponse{Status: http.StatusCreated, Data: []db.Task{newTask}})
	}).Methods("POST")

	r.HandleFunc("/api/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		//define the task Struct
		var editTask db.Task
		//make a decorder to read the body
		decoder := json.NewDecoder(r.Body)
		//decode the body into the edit task struct, if there's an error return
		if err := decoder.Decode(&editTask); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", err)
			return
		}
		//why do this???
		defer r.Body.Close()
		//get the id of the task, could also grab it off of the edit task struct
		id := mux.Vars(r)["id"]
		//try to update thte task TODO: update this task to not return string error, but sql.row and error
		_, err := db.UpdateTask(database, editTask, id)
		//if there's an error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", err)
			return
		}
		//return a response to the front
		utils.RespondWithJSON(w, http.StatusOK, taskResponse{Status: http.StatusOK, Data: []db.Task{editTask}})
	})
}
