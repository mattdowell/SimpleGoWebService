package obs

import (
	"GoWebService/src/basicdb/mgr"
	"fmt"
	"log"
)

/*
   Simple type struct for passing around data objects
*/
type SimpleDbType struct {
	Id     int32
	Name   string
	Number int32
}

//
// Inserts a simple type into the db test_table
// This is a method on the SimpleDbType struct
//
func (data *SimpleDbType) Insert(dbMgr *mgr.DBConn) error {
	db := dbMgr.Open()

	sqlStatement := `INSERT INTO test_table ("name", "number") VALUES ($1, $2) RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, data.Name, data.Number).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	dbMgr.Close()
	return err
}

//    Reads the given row ID into a simple struct and returns the struct
//
func Read(id_to_read int32, dbMgr *mgr.DBConn) SimpleDbType {
	db := dbMgr.Open()
	theReturn := SimpleDbType{}

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
	dbMgr.Close()
	return theReturn
}

//
// Writes the simple DB tpye to the database then returns it
// with a new ID
//
func Write(inType SimpleDbType, dbMgr *mgr.DBConn) SimpleDbType {
	db := dbMgr.Open()
	sqlStatement := "INSERT INTO test_table (name, number) VALUES ($1, $2) RETURNING id"
	id := 0
	// Do the insert and query for the ID
	err := db.QueryRow(sqlStatement, inType.Name, inType.Number).Scan(&id)
	if err != nil {
		panic(err)
	}
	return Read(int32(id), dbMgr)
}
