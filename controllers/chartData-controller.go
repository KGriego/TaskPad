package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Midlu/GoShit/Mine/TaskPad/db"
	"github.com/Midlu/GoShit/Mine/TaskPad/utils"
	"github.com/gorilla/mux"
)

type chartDataResponse struct {
	Status  int         `json:"status"`
	Data    []db.Report `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

//CreateChartDataRoutes make the routes for anything related to the tasks
func CreateChartDataRoutes(r *mux.Router, database *sql.DB) {
	//route for this one is: GET /api/chartdata/{id}
	r.HandleFunc("/api/chartdata/{id}", func(w http.ResponseWriter, r *http.Request) {
		//get the id of the task, could also grab it off of the edit task struct
		id := mux.Vars(r)["id"]
		rows, err := db.GetAllChartData(database, id)
		//there was an error trying to retrieve the records, stop here
		if err != nil {
			errorString := fmt.Sprintf("There was an error retrieving the records. Error: %s", err)
			utils.RespondWithError(w, http.StatusNotFound, errorString)
		}
		//return a response to the front
		utils.RespondWithJSON(w, http.StatusOK, chartDataResponse{Status: http.StatusOK, Data: rows})
	}).Methods("GET")
}
