package main


import (

	"fmt"

	"github.com/gin-gonic/gin"
 	"github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
)


type ContactUS struct {
	gorm.Model 
	Name string
	Email string
	Query string
	Message string
}


func main() {

	var db *gorm.DB 
	var err error


	db, err = gorm.// import "gorm.io/driver/mysql"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "root:Nabin&dangi320@tcp(127.0.0.1:3306)/nds_ndscontactus?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	


	fmt.Println("HERE")
}

