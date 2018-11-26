package main

import (
	//"bytes"
	//"context"
	"fmt"
	"net/http"
	//"database/sql"
	//"log"
	//"strconv"
	//"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"./models"
	"./db"
)

//var db, err = sql.Open("mysql", "yprima:Dbtjdwo1$$@tcp(14.63.127.152:10824)/battle_game")


func getTestString(c *gin.Context) {
	
	c.JSON(http.StatusOK , gin.H{ "result" : "testok"})
}

func getSendBoxList(c *gin.Context) {
	var sendbox models.SendBox 

	var sendboxs []models.SendBox
	
	//rows , err := db.Query("CALL SP_SEND_BOX_LIST(1);") //adhoc 쿼리로 변경 
	rows, err := db.Init().Query("SELECT boxidx , sender_address , boxmsg , send_wei FROM SendBox ORDER BY boxidx DESC;")

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

	var boxranking models.SendBoxRanking

	var boxrankings []models.SendBoxRanking

	//rows , err := db.Query("CALL SP_SEND_BOX_RANKING(1);")
	rows , err := db.Init().Query("SELECT 	sender_address , last_boxmsg , total_take_token	FROM SendBox_address ORDER BY total_take_token DESC LIMIT 5")
	
	if err != nil {
		fmt.Println(err.Error())
	}

	rank := 1
	
	for rows.Next() {
		err = rows.Scan(&boxranking.Sender_address , &boxranking.Last_boxmsg , &boxranking.Total_take_token)
		boxranking.Rank = rank
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

	row := db.Init().QueryRow("CALL SP_SEND_BOX_CAN_TAKE_SENDBOX(?,?,@o_return);" , boxidx , takeaddress)

	err := row.Scan(&outreturn.O_return)

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
func setRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/", getTestString)
		v1.GET("/sendaboxlist" , getSendBoxList)
		v1.GET("/sendaboxranking", getBoxRanking)
		v1.POST("/takeabox" , postTakeBox)
	}
	return router
  }


  
func main() {
	

	router := setRouter()
	
	fmt.Println("잘뜸?")
	router.Run(":3030")

}
