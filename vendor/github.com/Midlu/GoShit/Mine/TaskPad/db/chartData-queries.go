package db

import (
	"database/sql"
	"fmt"
)

//Chart structure for Chart
type Chart struct {
	ID     int    `json:"id,omitempty"`
	UserID string `json:"userid,omitempty"`
	Title  string `json:"title,omitempty"`
	Type   string `json:"type,omitempty"`
	Spec   spec   `json:"spec,omitempty"`
}

//GetAllChartData takes a db connection, and returns a tasks array and an error
func GetAllChartData(db *sql.DB, id string) ([]Report, error) {
	queryString := fmt.Sprintf("SELECT * FROM reports where id=%s;", id)
	//try to get the rows
	rows, err := db.Query(queryString)
	//there was an error trying to retrieve the records, stop here
	if err != nil {
		return nil, err
	}
	//why do they do this??
	defer rows.Close()
	//this says the we're making an array of unknown length with the struct of task
	reports := []Report{}
	//mapping over the row result?
	for rows.Next() {
		//declaring the task
		var r Report
		//copy the columns in the task struct, if there's an error doing this, stop here and return the error
		if err := rows.Scan(&r.ID, &r.UserID, &r.Title, &r.Type, &r.Spec); err != nil {
			return nil, err
		}
		//add the copied task into the array
		reports = append(reports, r)
	}
	//return the tasks
	return reports, nil
}
