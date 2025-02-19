package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 전역 DB 인스턴스

func ConnectDB() {
	/*
		dsn := "root:1234@tcp(localhost:3306)/twinkle?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("##### Failed to connect to database ##### : ", err)
		}

		fmt.Println("##### Database connected successfully ##### ")
		DB = db
	*/

	dsn := "root:1234@tcp(localhost:3306)/twinkle?charset=utf8mb4&parseTime=True&loc=Local"

	fmt.Println("##### Connecting to MySQL... #####") // ✅ MySQL 연결 시작 로그

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("##### Failed to connect to database #####")
		log.Fatal("Error details:", err) // ✅ MySQL 연결 실패 로그
	}

	fmt.Println("##### Database connected successfully ##### ") // ✅ 성공 로그
	DB = db
}
