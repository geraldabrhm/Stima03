package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	var database_config, uname, pass, host, port string
	fmt.Print("Database username: ")
	fmt.Scanln(&uname)
	fmt.Print("Database password: ")
	fmt.Scanln(&pass)
	fmt.Print("Database host(ex.127.0.0.0): ")
	fmt.Scanln(&host)
	fmt.Print("Database port: ")
	fmt.Scanln(&port)
	database_config = fmt.Sprintf("%s:%s@tcp(%s:%s)/random_scheme?charset=utf8mb4&parseTime=True&loc=Local", uname, pass, host, port)
	dsn := database_config
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Success connected to database")
}
