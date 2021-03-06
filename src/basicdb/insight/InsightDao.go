package insight

import (
	"GoWebService/src/basicdb/mgr"
	"fmt"
	"log"
)

/*
   Simple type struct for passing around data objects
*/
type InsightRecord struct {
	Id     int32
	Name   string
	Number int32
}

//    Inserts a simple type into the db test_table
// This is a method on the SimpleDbType struct
func (data *InsightRecord) Insert() error {
	db := mgr.Open()

	sqlStatement := `INSERT INTO test_table ("name", "number") VALUES ($1, $2) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, data.Name, data.Number).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	mgr.Close(db)
	return err;
}

/**
  Reads the given row ID into a simple struct and returns the struct
*/
func Read(id_to_read int32) InsightRecord {
	db := mgr.Open()
	theReturn := InsightRecord{}

	sqlRead := fmt.Sprintf("select id, name, number from test_table where id = %d", id_to_read)

	rows, err := db.Query(sqlRead)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&theReturn.Id, &theReturn.Name, &theReturn.Number)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	mgr.Close(db)
	return theReturn
}
