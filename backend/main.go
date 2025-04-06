// main.go
package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/controllers"
	"github.com/johnnydev/kids-cafe-backend/firebase"
	// "github.com/johnnydev/kids-cafe-backend/models"
)

func main() {

	fmt.Print("✅✅ 터미널 확인")

	// Firebase 초기화
	firebase.InitializeFirebase() // Firebase 초기화 함수 호출

	// config.ConnectDB()

	// if config.DB == nil {
	// 	log.Fatal("❌ Database connection failed!") // DB 연결 실패 체크
	// } else {
	// 	fmt.Println("✅ Database is connected and ready!") // DB 정상 연결 확인
	// }

	// r = router
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Kids Cafe API!"})
	})

	// CORS 미들웨어 설정
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"}, // 허용할 출처
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	r.GET("/api/branches", controllers.GetBranches)
	r.GET("/api/partyrooms/:branch_id", controllers.GetPartyrooms)
	r.GET("/api/selectedroom/:room_id", controllers.GetSelectedRoom)
	r.POST("/api/select-time", controllers.SaveSelectTime)
	r.GET("/api/foodoptions", controllers.Getfoodoptions)
	r.GET("/api/save-party", controllers.SaveParty)
	r.GET("/api/cafe-menu", controllers.GetMenu)
	// 엔드포인트 확인
	/*
		r.GET("/api/branches", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "왜 안되는 거야 Branch data fetched successfully!",
			})
		})
	*/

	// CORS 미들웨어 설정
	r.Use(cors.Default()) // 기본 설정 (모든 도메인 허용)

	r.Run(":8081") // 백엔드 서버 실행

	r.SetTrustedProxies(nil) // 프록시를 신뢰하지 않도록 설정

}
