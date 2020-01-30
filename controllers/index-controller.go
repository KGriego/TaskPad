package controllers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

//SpaHandler struct for the single page app
type SpaHandler struct {
	StaticPath string
	IndexPath  string
}

func (h SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("At SpaHandler route.")
	//get the path from the url???? but whyyyyy?? to access the directory/files from where the server is located
	path, err := filepath.Abs(r.URL.Path)
	//if there is an err stop here
	if err != nil {
		//respond with an error to the request
		http.Error(w, err.Error(), http.StatusBadRequest)
		//does the fatal stop the server? If so, do I still need the return statement?
		log.Fatalln("There was an error: ", err.Error(), http.StatusInternalServerError)
		return
	}
	//join the file path to locate the spa files
	path = filepath.Join(h.StaticPath, path)
	//try to find the path
	_, err = os.Stat(path)
	//if the directory does not exist stop here
	if os.IsNotExist(err) {
		//path doesn't exist, serve the index file? makes no sense to me
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	} else if err != nil {
		//if a different error is returned
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln("There was an error: ", err.Error(), http.StatusInternalServerError)
		return
	}
	//otherwise just serve the default file?
	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}
