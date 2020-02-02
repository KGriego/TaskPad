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
	dbConnectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	databaseURL := os.Getenv("DATABASE_URL")
	//if no db connection string exists stop here
	if dbConnectionString == "" && databaseURL == "" {
		log.Fatalln("A DB connection string is required. Aborting")
		return nil, errors.New("A DB connection string is required. Aborting")
	} else if databaseURL != "" {
		//attemp to connect to db with accquired string
		db, err := sql.Open("postgres", databaseURL)
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
	} else {
		//attemp to connect to db with accquired string
		db, err := sql.Open("postgres", dbConnectionString)
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
}
