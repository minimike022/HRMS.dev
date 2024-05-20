package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func Connect() *sql.DB{
	dbConn, err := sql.Open("mysql","root:@tcp(localhost:3306)/recruitmentms")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database Connected")
	return dbConn
}