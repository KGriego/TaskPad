package db

import (
	"database/sql"
)

//UserTag structure for tasks
type UserTag struct {
	UserID string `json:"userid"`
	Tag    string `json:"title"`
}

//GetUserTags takes a db connection, and returns a tasks array and an error
func GetUserTags(db *sql.DB) ([]UserTag, error) {
	queryString := "SELECT * FROM usertags;"
	//try to get the rows
	rows, err := db.Query(queryString)
	//there was an error trying to retrieve the records, stop here
	if err != nil {
		return nil, err
	}
	//why do they do this??
	defer rows.Close()
	//this says the we're making an array of unknown length with the struct of task
	usertags := []UserTag{}
	//mapping over the row result?
	for rows.Next() {
		//declaring the task
		var ut UserTag
		//copy the columns in the task struct, if there's an error doing this, stop here and return the error
		if err := rows.Scan(&ut.UserID, &ut.Tag); err != nil {
			return nil, err
		}
		//add the copied task into the array
		usertags = append(usertags, ut)
	}
	//return the tasks
	return usertags, nil
}
