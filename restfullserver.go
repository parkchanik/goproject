package main

import (
	"fmt"
	"net/http"
	"database/sql"
	//"log"
	//"strconv"
	//"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "yprima:Dbtjdwo1$$@tcp(14.63.127.152:10824)/battle_game")
 

type User struct {
	User_no 		string 
	Server_idx 		int 
	Publisher_id 	string 
	Auth_idx 		int 
	Nickname 		string 
	Gold 			int 
	Input_time 		string

}


type Person struct {
	Id         int
	First_Name string
	Last_Name  string
}

func getById(c *gin.Context) {
	var user User 

	id := c.Param("id")
	//row := db.QueryRow("select id, username, first_name, middle_name, last_name, email, mobile_phone, login_attempt, active_status FROM MEMBER_INFO where id = ?;", id)
	row := db.QueryRow("select user_no , server_idx , publisher_id , auth_idx , nickname , gold , input_time FROM MEMBER_INFO where user_no = ?;", id)
	
	err = row.Scan(&user.User_no, &user.Server_idx, &user.Publisher_id, &user.Auth_idx, &user.Nickname, &user.Gold, &user.Input_time)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, user)
	}
}


func main() {
	

	router := gin.Default()
	router.GET("/api/user/:id", getById)
	
	router.GET("/person/:id", func(c *gin.Context) {
		var (
			person Person
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select id, first_name, last_name from person where id = ?;", id)
		err = row.Scan(&person.Id, &person.First_Name, &person.Last_Name)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})
	fmt.Println("test")
	router.Run(":8000")

}