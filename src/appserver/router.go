package appserver

import (
	"GoWebService/src/basicdb/obs"
	"GoWebService/src/json"
	"fmt"
	"net/http"
)

// Handles the observation read
func HandleReadObs() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		theType := obs.SimpleDbType{}
		theType = obs.Read(1)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", json.Marshall(theType))
	}
}

// Writes a new record to the DB based on the JSON content of the body
func HandleWriteObs() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	}
}
