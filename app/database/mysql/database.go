package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)



func Connect() *sql.DB{
	db, err := sql.Open("mysql","root:@tcp(localhost:3306)/recruitmentms")
	
	if err != nil {
		panic(err.Error())
	}
	fmt.Print("connected to database")

	return db
}

