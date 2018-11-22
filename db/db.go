package db

import (
	_ "github.com/go-sql-driver/mysql"
  	"database/sql"
	"fmt"
)

func Init() (*sql.DB) {
 
	
	db, err := sql.Open("mysql", "dbchanik:1q2w3e$r@tcp(localhost:3306)/aboxdb")
 
  	checkErr(err)

	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	checkErr(err)
	fmt.Printf("Connection successfully")

	return db
}

func checkErr(err error) {
  if err != nil {
    fmt.Print(err.Error())
  }
}