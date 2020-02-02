package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Midlu/TaskPad/db"
	"github.com/Midlu/TaskPad/utils"
)

type taskMetricResponse struct {
	Status int             `json:"status"`
	Data   []db.TaskMetric `json:"data"`
	Error  string          `json:"error"`
}

//CreateTaskMetricRoutes make the routes for anything related to the tasks
func CreateTaskMetricRoutes(r *mux.Router, database *sql.DB) {
	r.HandleFunc("/api/taskmetrics/daily", func(w http.ResponseWriter, r *http.Request) {
		log.Println("at taskmetrics route")
		vars := r.URL.Query()
		rows, err := db.GetTaskMetrics(database, vars)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Fatalln("There was an error retrieving the records", err)
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, taskMetricResponse{Status: http.StatusOK, Data: rows})
		// rows, err := db.GetPendingTasks(database, vars)
		//there was an error trying to retrieve the records, stop here
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	log.Fatalln("There was an error retrieving the records", err)
		// 	return
		// }
		//return a response to the front
		// utils.RespondWithJSON(w, http.StatusOK, taskMetricResponse{Status: http.StatusOK, Data: rows})
	}).Methods("GET")
}
