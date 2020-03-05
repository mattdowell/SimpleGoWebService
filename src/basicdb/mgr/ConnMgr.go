package mgr

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)


/*
 https://tour.golang.org/methods/1
 Go does not have classes. However, you can define methods on types.
 A method is a function with a special receiver argument.
 The receiver appears in its own argument list between the func keyword and the method name.
  func (v DBConn) open() *sql.DB
 Remember: a method is just a function with a receiver argument.
 You can declare a method on non-struct types, too.
 */
type DBConn struct {
	Host string
	Port int
	User string
	Password string
	Dbname  string
	connection *sql.DB
}

/*
  Opens the database connection and returns a pointer to it
*/
func (p *DBConn) Open()  *sql.DB  {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	p.connection = db;
	return db
}

/*
 Closes the database connection
*/
func (db *DBConn) Close() {
	if db.connection != nil {
		defer db.connection.Close()
	}
}
