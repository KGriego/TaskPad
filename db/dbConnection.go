package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	//need this to connect to db

	_ "github.com/lib/pq"
)

//Connection func to connect to db
func Connection() (*sql.DB, error) {
	log.Println("Attempting to connect to DB")
	//attemp to get db connection string from env
	env := os.Getenv("ENV")
	//if no db connection string exists stop here
	if env == "" {
		log.Fatalln("A DB connection string is required. Aborting")
		return nil, errors.New("A DB connection string is required. Aborting")
	} else if env == "PROD" {
		databaseURL := os.Getenv("POSTGRES_CONNECTION_STRING_PROD")
		log.Println("In Prod")
		return testConnection(databaseURL)
	} else {
		log.Println("In Dev")
		dbConnectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
		return testConnection(dbConnectionString)
	}
}

func testConnection(dbString string) (*sql.DB, error) {
	//attemp to connect to db with accquired string
	db, err := sql.Open("postgres", dbString)
	//if there is an err when opening the connection, stop here
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	//ping the db, if something returns, connection was not fully successful
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	//connection was successful, return db connection and no error
	log.Println("Connection successful")
	return db, nil
}
