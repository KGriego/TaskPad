package db

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
)

//Report structure for tasks
type Report struct {
	ID     int    `json:"id,omitempty"`
	UserID string `json:"userid,omitempty"`
	Title  string `json:"title,omitempty"`
	Type   string `json:"type,omitempty"`
	Spec   spec   `json:"spec,omitempty"`
}

type spec struct {
	Tags   []string   `json:"tags,omitempty"`
	Groups []pieGroup `json:"groups,omitempty"`
}

type pieGroup struct {
	Name string `json:"name,omitempty"`
	Spec string `json:"spec,omitempty"`
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (s *spec) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &s)
}

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (s spec) Value() (driver.Value, error) {
	return json.Marshal(s)
}

//GetAllReports takes a db connection, and returns a tasks array and an error
func GetAllReports(db *sql.DB) ([]Report, error) {
	queryString := "SELECT * FROM reports;"
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

//DeleteSpecificReport takes a db connection, and returns a tasks array and an error
func DeleteSpecificReport(db *sql.DB, id string) error {
	//try to get the rows
	_, err := db.Exec("DELETE FROM reports WHERE id=$1;", id)
	//return the tasks
	return err
}

//CreateReport takes a db connection, and returns a tasks array and an error
func CreateReport(db *sql.DB, r Report) error {
	log.Println(r)
	//try to insert the report into the db, if there's an error, return the error
	if err := db.QueryRow(
		"INSERT INTO reports(userid, title, type, spec) VALUES($1, $2, $3, $4) RETURNING id",
		r.UserID, r.Title, r.Type, r.Spec).Scan(&r.ID); err != nil {
		return err
	}
	return nil
}
