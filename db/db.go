package db

import (
	_ "github.com/go-sql-driver/mysql"
  	"database/sql"
	"fmt"
)

func Init() (*sql.DB) {
 
	
	//db, err := sql.Open("mysql", "dbchanik:1q2w3e$r@tcp(localhost:3306)/aboxdb")
	db, err := sql.Open("mysql", "dbchanik:1q2w3e$r@tcp(tivasdbinstance.c1bdzesf6m8r.ap-northeast-2.rds.amazonaws.com:3306)/aboxdb")
	
  checkErr(err)
	fmt.Printf("ccccccc")
	//defer db.Close()
  // defer 로 클로즈 하면 안된다?
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