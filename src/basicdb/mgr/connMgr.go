package mgr

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "navio"
	password = ""
	dbname   = "navio"
)

/*
  Opens the database connection and returns a pointer to it
*/
func Open()  *sql.DB  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	return db
}

/*
 Closes the database connection
*/
func Close(db *sql.DB) {
	defer db.Close()
}
