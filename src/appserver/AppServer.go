package appserver

import (
	"GoWebService/src/basicdb/mgr"
	"log"
	"net/http"
)

// Validation of record
// Store a blob
// Change DB to mongo or some other no-sql
func ServerMain(config *mgr.DBConn) {

	AddGorillaRoutes(config)

	log.Println("---Get Ready---")
	log.Fatal(http.ListenAndServe(":8080", nil))
}


