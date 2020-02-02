package controllers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

//AddAPIRouterHandler add api routes here to help keep main.go file clean
func AddAPIRouterHandler(r *mux.Router, database *sql.DB) {
	//call the CreateTaskRoutes from task-controller to help keep this file clean as well
	CreateTaskRoutes(r, database)
	CreateReportRoutes(r, database)
	CreateUserTagRoutes(r, database)
	CreateTaskMetricRoutes(r, database)
	CreateChartDataRoutes(r, database)
	//make this last to hit the api routes first.
	//define the index route for the SPA files
	spa := SpaHandler{StaticPath: "ui-dist", IndexPath: "index.html"}
	//make the home route for the SPA
	r.PathPrefix("/").Handler(spa)
}

//MAYBE ONE DAY, UNTIL THEN COMMENTED OUT
//APIRoutesHandlerFunc definition
// type APIRoutesHandlerFunc struct {
// 	http.Handler
// }

//APIRoutesHandlerFunc definition
//Want to do it this way, don't know how though.
// func (h APIRoutesHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Println("APIRoutesHandlerFunc: Adding API Routers to router.")
// }
