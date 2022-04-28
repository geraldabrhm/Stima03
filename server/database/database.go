package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	const MYSQL = "root:asdf@tcp(127.0.0.1:3306)/penyakit?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// const POSTGRES = "host=localhost user=postgres password=root dbname=dna_tracker port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// dsn := POSTGRES
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to database")
	}
	fmt.Println("Success connected to database")
}
