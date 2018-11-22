package main

import (
	//"bytes"
	//"context"
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
	Sender_address 	string 
	Boxmsg 			string 
	Send_wei 		int 
}

type SendBoxRanking struct {
	Rank 				int
	Sender_address 		string
    Last_boxmsg			string
	Total_take_token 	int
			
}

type Person struct {
	Id         int
	First_Name string
	Last_Name  string
}

func getTestString(c *gin.Context) {
	
	c.JSON(http.StatusOK , gin.H{ "result" : "testok"})
}

func getSendBoxList(c *gin.Context) {
	var sendbox SendBox 

	var sendboxs []SendBox
	
	//rows , err := db.Query("CALL SP_SEND_BOX_LIST(1);") //adhoc 쿼리로 변경 
	rows, err := db.Query("SELECT boxidx , sender_address , boxmsg , send_wei FROM SendBox ORDER BY boxidx DESC;")

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		err = rows.Scan(&sendbox.Boxidx, &sendbox.Sender_address, &sendbox.Boxmsg, &sendbox.Send_wei)
		sendboxs = append(sendboxs, sendbox)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()

	
	c.Header("Access-Control-Allow-Origin" , "*")

	c.JSON(http.StatusOK, gin.H{
		"result": sendboxs,
		"count":  len(sendboxs),
	})
	
	
	
}

func getBoxRanking(c *gin.Context) {

	var boxranking SendBoxRanking

	var boxrankings []SendBoxRanking

	//rows , err := db.Query("CALL SP_SEND_BOX_RANKING(1);")
	rows , err := db.Query("SELECT 	sender_address , last_boxmsg , total_take_token	FROM SendBox_address ORDER BY total_take_token DESC LIMIT 5")
	
	if err != nil {
		fmt.Println(err.Error())
	}

	rank := 1
	
	for rows.Next() {
		err = rows.Scan(&rank , &boxranking.Sender_address , &boxranking.Last_boxmsg , &boxranking.Total_take_token)
		boxrankings = append(boxrankings , boxranking)

		if err != nil {
			fmt.Println(err.Error())
		}

		rank++
	}

	defer rows.Close()

	c.Header("Access-Control-Allow-Origin" , "*")

	c.JSON(http.StatusOK, gin.H{
		"result": boxrankings,
		"count":  len(boxrankings),
	})

}

func postTakeBox(c *gin.Context) {
		
		
	type OutReturn struct {
		O_return         int
	}

	var outreturn OutReturn

	boxidx := c.PostForm("boxidx")
	takeaddress := c.PostForm("takeaddress")

	row := db.QueryRow("CALL SP_SEND_BOX_CAN_TAKE_SENDBOX(?,?,@o_return);" , boxidx , takeaddress)

	err = row.Scan(&outreturn.O_return)

	c.Header("Access-Control-Allow-Origin" , "*")
	c.JSON(http.StatusOK, gin.H{
		"result": outreturn ,
		"error" : err ,
	})


}

/*

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

*/
func main() {
	

	router := gin.Default()
	//router.GET("/api/user/:user_no", getMemberInfoByUserNo)
    router.GET("/" , getTestString)
	router.GET("/sendaboxlist", getSendBoxList)
	router.GET("/sendaboxranking", getBoxRanking)
	router.POST("/takeabox" , postTakeBox)
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
	fmt.Println("잘뜸?")
	router.Run(":3030")

}