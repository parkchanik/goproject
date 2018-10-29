package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"log"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql/driver/mysql"
)

func main() {
	http.HandleFunc("/" , func(w http.ResponseWriter , r *http.Request) {
		fmt.Fprintln(w , "welcome!")
	})

	http.HandleFunc("/about" , func( w http.ResponseWriter , r *http.Request) {
		fmt.Fprintln(w , "about")
	})

	http.ListenAndServe(":8000" , nil)

}