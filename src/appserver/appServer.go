package appserver

import (
	"log"
	"net/http"
)

// Read and parse JSON into DB records
// Validation of record
// Store a blob
// Change DB to mongo or some other no-sql
func AppServerMain() {


	http.HandleFunc("/read", HandleReadObs())
	http.HandleFunc("/write", HandleWriteObs())

	log.Println("---Get Ready---")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
