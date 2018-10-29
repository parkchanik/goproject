package main

import (
	"bytes"
	"fmt"
	"net/http"
	"database/sql"
	//"log"
	//"strconv"
	//"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//var db, err = sql.Open("mysql", "yprima:Dbtjdwo1$$@tcp(14.63.127.152:10824)/battle_game")
var db, err = sql.Open("mysql", "dbchanik:1q2w3e$r@tcp(localhost:3306)/aboxdb")
 

type User struct {
	User_no 		string 
	Server_idx 		int 
	Publisher_id 	string 
	Auth_idx 		int 
	Nickname 		string 
	Gold 			int 
	Input_time 		string

}

type SendBox struct {
	Boxidx 			int 
	Sender_address 		string 
	Boxmsg 			string 
	Send_wei 		int 
}


type Person struct {
	Id         int
	First_Name string
	Last_Name  string
}

func getMemberInfoByUserNo(c *gin.Context) {
	var user User 

	user_no := c.Param("user_no")
	//row := db.QueryRow("select id, username, first_name, middle_name, last_name, email, mobile_phone, login_attempt, active_status FROM MEMBER_INFO where id = ?;", id)
	row := db.QueryRow("select user_no , server_idx , publisher_id , auth_idx , nickname , gold , input_time FROM MEMBER_INFO where user_no = ?;", user_no)
	
	err = row.Scan(&user.User_no, &user.Server_idx, &user.Publisher_id, &user.Auth_idx, &user.Nickname, &user.Gold, &user.Input_time)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func getSendBoxList(c *gin.Context) {
	var sendbox SendBox 

	//var sendboxs []SendBox
	//row := db.QueryRow("select id, username, first_name, middle_name, last_name, email, mobile_phone, login_attempt, active_status FROM MEMBER_INFO where id = ?;", id)
	row := db.QueryRow("CALL SP_SEND_BOX_LIST(1);")
	/*
	for rows.Next() {
		err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name)
		persons = append(persons, person)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
*/
	err = row.Scan(&sendbox.Boxidx, &sendbox.Sender_address, &sendbox.Boxmsg, &sendbox.Send_wei)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, sendbox)
	}
}



func putMemberInfoByUserNo(c *gin.Context) {

	var buffer bytes.Buffer
	user_no := c.Query("user_no")

	gold := c.PostForm("gold")

	stmt , err := db.Prepare("UPDATE MEMBER_INFO SET gold = ? WHERE user_no = ?;")

	if err != nil {
		fmt.Print(err.Error())
	}

	_ , err = stmt.Exec(gold , user_no)

	if err != nil {
		fmt.Print(err.Error())
	}

	buffer.WriteString(gold)
	buffer.WriteString(" ")
	buffer.WriteString(user_no)

	defer stmt.Close()

	name := buffer.String()

	c.JSON(http.StatusOK , name)



}


func main() {
	

	router := gin.Default()
	router.GET("/api/user/:user_no", getMemberInfoByUserNo)

	router.GET("/sendboxlist", getSendBoxList)
	
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