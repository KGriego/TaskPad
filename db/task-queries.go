package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/lib/pq"

	"github.com/Midlu/GoShit/Mine/TaskPad/utils"
)

//Task structure for tasks
type Task struct {
	ID        int      `json:"id"`
	UserID    string   `json:"userid"`
	Title     string   `json:"title"`
	Due       string   `json:"due"`
	Completed bool     `json:"completed"`
	Effort    float64  `json:"effort"`
	Tags      []string `json:"tags"`
	Notes     string   `json:"notes"`
}

//makeArrayToStringForTags makes the string format to talk with postgres db
func makeArrayToStringForTags(t Task) string {
	//for every string in the tags array
	for i, s := range t.Tags {
		//wrap it in double quotes?
		t.Tags[i] = strconv.Quote(s)
	}
	//merge the string and return it
	return "{" + strings.Join(t.Tags, ",") + "}"
}

//GetPendingTasks takes a db connection, and returns a tasks array and an error
func GetPendingTasks(db *sql.DB, urlParams url.Values) ([]Task, error) {
	var queryString string
	if urlParams.Get("pending") != "" {
		//if the tasks are not completed
		queryString = fmt.Sprintf("SELECT * FROM tasks WHERE completed != %s;", urlParams.Get("pending"))
	} else if urlParams.Get("from") != "" {
		//parse and format the from and to date
		from, to, err := utils.GetToAndFrom(urlParams)
		//there was an error trying to retrieve format the dates, stop here
		if err != nil {
			return nil, err
		}
		//add the from and to into the query string
		queryString = fmt.Sprintf("SELECT * FROM tasks WHERE due >= '%s' AND due <= '%s';", from, to)
	} else {
		//return no data and no params were made
		//this says the we're making an array of unknown length with the struct of task
		tasks := []Task{}
		return tasks, nil
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
	tasks := []Task{}
	//mapping over the row result?
	for rows.Next() {
		//declaring the task
		var t Task
		//copy the columns in the task struct, if there's an error doing this, stop here and return the error
		//pq.Array somehow converts it from []int8 to []string, works for me though
		if err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Due, &t.Completed, &t.Effort, pq.Array(&t.Tags), &t.Notes); err != nil {
			return nil, err
		}
		//add the copied task into the array
		tasks = append(tasks, t)
	}
	//return the tasks
	return tasks, nil
}

//CreateTask function to insert task into db
func CreateTask(db *sql.DB, t Task) error {
	//format the tags array to be able to insert the task into db
	tags := makeArrayToStringForTags(t)
	err := db.QueryRow(
		"INSERT INTO tasks(userid, title, due, completed, effort, tags, notes) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		t.UserID, t.Title, t.Due, t.Completed, t.Effort, tags, t.Notes).Scan(&t.ID)
	if err != nil {
		return err
	}
	return nil
}

//UpdateTask function to udpate a task
func UpdateTask(db *sql.DB, t Task, id string) (string, error) {
	//format the tags array to be able to update the task in db
	tags := makeArrayToStringForTags(t)
	_, err := db.Exec("UPDATE tasks SET userid=$1, title=$2, due=$3, completed=$4, effort=$5, tags=$6, notes=$7 WHERE id=$8", t.UserID, t.Title, t.Due, t.Completed, t.Effort, tags, t.Notes, id)
	if err != nil {
		return "", err
	}
	return "", nil
}
