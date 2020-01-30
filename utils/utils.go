package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//GetToAndFrom gets and formats time for db query string
func GetToAndFrom(urlParams url.Values) (string, string, error) {
	//parse the from date
	fromParse, err := time.Parse(time.RFC3339, urlParams.Get("from"))
	//if theres an error parse from date
	if err != nil {
		return "", "", err
	}
	//parse the to date
	toParse, err := time.Parse(time.RFC3339, urlParams.Get("to"))
	//if there an error parsing the to date
	if err != nil {
		return "", "", err
	}
	//from the UTC part from the end of the string since db query doesn't like that, I think
	from := strings.TrimSuffix(fromParse.String(), " UTC")
	to := strings.TrimSuffix(toParse.String(), " UTC")
	//return the strings and no error
	return from, to, nil
}

//RespondWithJSON util func to help respond easier TODO: Add error message to return to front
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//make the payload in a json format
	response, err := json.Marshal(payload)
	//if there's an error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("There was an error retrieving the records", err)
		return
	}
	//respond to the front with the json response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

//RespondWithError func to help send error messages easier respond easier
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
