package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	c "github.com/Midlu/GoShit/Mine/TaskPad/controllers"
	"github.com/Midlu/GoShit/Mine/TaskPad/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//init runs before main, used to fetch env variable, just in case
func init() {
	//try to get the .env file, return an error if not found
	err := godotenv.Load()
	//if there is an error
	if err != nil {
		log.Fatalln("No .env file found! Please make one")
	}
}

func main() {
	db, err := db.Connection()
	//if there is an error, is there a better way to handle errors? Like can I throw it?
	if err != nil {
		log.Fatalln("There was an error. Stopping server", err)
		return
	}
	//defer is telling it to wait until functions nearby finish, but why this early?
	defer db.Close()
	//create the server to be able to customize the handles instead of it being defaulted
	router := mux.NewRouter()
	//create another route handle for the api calls
	log.Println("Going to add routes now")
	//pass the router, and make the routes here
	c.AddAPIRouterHandler(router, db)
	//don't know how to implement this way. One day I will
	// router.PathPrefix("/api/").Handler(c.APIRoutesHandlerFunc)
	log.Println("Added routes")
	//get the port from the env variables
	port := os.Getenv("PORT")
	//if there is no port set a port
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	//create server options
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//create a string to print the console to know what port is being used
	log.Println(fmt.Sprintf("🌎 is listening on: %s", port))
	log.Fatal(srv.ListenAndServe())
}
