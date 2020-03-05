package appserver

import (
	"GoWebService/src/basicdb/mgr"
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

var dbMgr *mgr.DBConn

// Builds all the routes
// http://www.gorillatoolkit.org/pkg/mux
//
func AddGorillaRoutes(inDbConn *mgr.DBConn) {
	dbMgr = inDbConn
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
	theType = obs.Read(getNumber(r), dbMgr)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", json.Marshall(theType))
}

//
// Get the object number from the path
//
func getNumber(r *http.Request) int32 {
	vars := mux.Vars(r)
	idStr := vars["id"]

	i1, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Error converting: " + idStr)
		return 1
	}
	// Force to int32
	return int32(i1)
}

// Writes a new record to the DB based on the JSON content of the body
func HandleWriteObs(w http.ResponseWriter, r *http.Request) {
	// Create and write the Object
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	convertedType := json.UnMarshall(bodyString)
	savedType := obs.Write(convertedType, dbMgr)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// Send the JSON back with the new ID
	fmt.Fprintf(w, "%s", json.Marshall(savedType))
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
