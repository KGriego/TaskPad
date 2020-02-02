package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/Midlu/GoShit/Mine/TaskPad/utils"
)

//TaskMetric strucuture for task metrics
type TaskMetric struct {
	Day    time.Time `json:"day"`
	Effort int       `json:"effort"`
}

//GetTaskMetrics takes a db connection, and returns a tasks array and an error
func GetTaskMetrics(db *sql.DB, urlParams url.Values) ([]TaskMetric, error) {
	//parse and format the from and to date
	user := urlParams.Get("user")
	from, to, err := utils.GetToAndFrom(urlParams)
	//format the query string
	queryString := fmt.Sprintf("SELECT date(due), effort FROM tasks WHERE (due between '%s' and '%s') AND userid = '%s' GROUP BY date(due), effort ORDER BY date(due)", from, to, user)
	//there was an error trying to retrieve the records, stop here
	if err != nil {
		return nil, err
	}
	//try to get the rows
	rows, err := db.Query(queryString)
	//there was an error trying to retrieve the records, stop here
	if err != nil {
		return nil, err
	}
	//why do they do this??
	defer rows.Close()
	//this says the we're making an array of unknown length with the struct of task
	taskMetrics := []TaskMetric{}
	//mapping over the row result?
	for rows.Next() {
		//declaring the task
		var t TaskMetric
		//copy the columns in the task struct, if there's an error doing this, stop here and return the error
		//pq.Array somehow converts it from []int8 to []string, works for me though
		if err := rows.Scan(&t.Day, &t.Effort); err != nil {
			return nil, err
		}
		//add the copied task into the array
		taskMetrics = append(taskMetrics, t)
	}
	//return the tasks
	return taskMetrics, nil
}
