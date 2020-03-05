package main

import (
	"GoWebService/src/appserver"
	"GoWebService/src/props"
	"GoWebService/src/basicdb/mgr"
	"log"
	"strconv"
)

/*
POINTERS!
https://tour.golang.org/moretypes/1
---------
The & operator generates a pointer to its operand.
i := 42
p = &i

The * operator denotes the pointer's underlying value.
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p

 To access the field X of a struct when we have the struct pointer
 p we could write (*p).X. However, that notation is cumbersome, so
 the language permits us instead to write just p.X

 */
func main() {


	p, err := props.ReadPropertiesFile("src/props/database.properties")
	if err != nil {
		panic(err)
	}
	params := mgr.DBConn{}
	params.Host = p["host"]
	params.Dbname = p["dbname"]
	params.User = p["user"]
	params.Port, _ = strconv.Atoi(p["port"])
	params.Password = p["password"]

	appserver.ServerMain(&params)

	log.Println("Using Database: " + p["database"])

}
