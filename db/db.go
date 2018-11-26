package db

import (
	_ "github.com/go-sql-driver/mysql"
  	"database/sql"
	"fmt"
)

func Init() (*sql.DB) {
 
	
	//db, err := sql.Open("mysql", "dbchanik:1q2w3e$r@tcp(localhost:3306)/aboxdb")
	db, err := sql.Open("mysql", "b7b0d77b49fd3a:5e9c3426@tcp(us-cdbr-iron-east-04.cleardb.net:3306)/heroku_349e294b225dd6c")
	
  checkErr(err)

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