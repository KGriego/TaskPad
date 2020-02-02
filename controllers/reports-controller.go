package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Midlu/TaskPad/db"
	"github.com/Midlu/TaskPad/utils"
	"github.com/gorilla/mux"
)

type reportsResponse struct {
	Status  int         `json:"status"`
	Data    []db.Report `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

//CreateReportRoutes make the routes for anything related to the tasks
func CreateReportRoutes(r *mux.Router, database *sql.DB) {
	//route for this one is: DELETE /api/reports/id
	r.HandleFunc("/api/reports/{id}", func(w http.ResponseWriter, r *http.Request) {
		//get the id of the task, could also grab it off of the edit task struct
		id := mux.Vars(r)["id"]
		if err := db.DeleteSpecificReport(database, id); err != nil {
			errorString := fmt.Sprintf("There was an error retrieving the records. Error: %s", err)
			utils.RespondWithError(w, http.StatusInternalServerError, errorString)
		}
		//return a response to the front
		utils.RespondWithJSON(w, http.StatusOK, reportsResponse{Status: http.StatusOK, Message: "Success"})
	}).Methods("DELETE")

	//route for this one is: POST /api/reports
	r.HandleFunc("/api/reports", func(w http.ResponseWriter, r *http.Request) {
		log.Println("at POST api/reports, creating report")
		//define the task Struct
		var newReport db.Report
		//make a decorder to read the body
		decoder := json.NewDecoder(r.Body)
		//decode the body into the new task struct, if there's an error return
		if err := decoder.Decode(&newReport); err != nil {
			errorString := fmt.Sprintf("There was an error decoding the body. Error: %s", err)
			log.Fatalln(errorString)
			utils.RespondWithError(w, http.StatusInternalServerError, errorString)
		}
		//why do this???
		defer r.Body.Close()
		//try to add the new report to the db
		err := db.CreateReport(database, newReport)
		//there was an error trying to retrieve the records, stop here
		if err != nil {
			errorString := fmt.Sprintf("There was an error creating the record. Error: %s", err)
			log.Fatalln(errorString)
			utils.RespondWithError(w, http.StatusInternalServerError, errorString)
		}
		//return a response to the front
		utils.RespondWithJSON(w, http.StatusOK, reportsResponse{Status: http.StatusOK, Message: "Success"})
		//update headers and return this
		log.Println("Finished creating report.")
	}).Methods("POST")

	//route for this one is: GET /api/reports
	r.HandleFunc("/api/reports", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.GetAllReports(database)
		//there was an error trying to retrieve the records, stop here
		if err != nil {
			errorString := fmt.Sprintf("There was an error retrieving the records. Error: %s", err)
			utils.RespondWithError(w, http.StatusNotFound, errorString)
		}
		//return a response to the front
		utils.RespondWithJSON(w, http.StatusOK, reportsResponse{Status: http.StatusOK, Data: rows})
	}).Methods("GET")
}
