package appserver

import (
	"GoWebService/src/basicdb/obs"
	"GoWebService/src/json"
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

//func HandleFunc(pattern string, f func(ResponseWriter, *Request)) {
//	DefaultServeMux.HandleFunc(pattern, f)
//}
// Builds all the routes
func AddRoutes() {
	// build the routes
	http.HandleFunc("/read/{id}", HandleReadObs)
	http.HandleFunc("/write", HandleWriteObs)
}


// Builds all the routes
// http://www.gorillatoolkit.org/pkg/mux
//
func AddGorillaRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/read/{id}", HandleReadObs).Methods("GET")
	r.HandleFunc("/write", HandleWriteObs).Methods("POST")
	http.Handle("/", r)
}

//
// Handles the observation read
//
func HandleReadObs(w http.ResponseWriter, r *http.Request) {
		theType := obs.SimpleDbType{}
		theType = obs.Read(getNumber(r))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", json.Marshall(theType))
}

//
// Get the object number from the path
//
func getNumber(r *http.Request) int32 {
	path := r.URL.Path
	log.Println("Path: " + path)
	i1, err := strconv.Atoi(path)
	if err == nil {
		fmt.Println(i1)
		return 1
	}

	// Force to int32
	return int32(i1)
}

// Writes a new record to the DB based on the JSON content of the body
func HandleWriteObs(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		// Create and write the Object
		// Send the JSON back with the new ID
}

// Writes an image
func Image(w http.ResponseWriter, r *http.Request) {
	// Open a JPG file.
	f, _ := os.Open("/home/sam/coin.jpg")

	// Read the entire JPG file into memory.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Set the Content Type header.
	w.Header().Set("Content-Type", "image/jpeg")

	// Write the image to the response.
	w.Write(content)
}
