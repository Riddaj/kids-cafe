// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // ✅ .env 파일 로딩용 패키지 추가

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/controllers"
	"github.com/johnnydev/kids-cafe-backend/firebase"
	"github.com/johnnydev/kids-cafe-backend/models"
	// "github.com/johnnydev/kids-cafe-backend/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env 파일을 불러오지 못했습니다 (배포 환경일 수 있음)")
	}

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

	// ✅ Render용 Health Check
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// CORS 미들웨어 설정
	//Authorization 추가
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://kids-cafe-booking-project.web.app",
			"http://localhost:8080", //로컬테스트용
		}, // 허용할 출처
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept",
			"Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/api/branches", controllers.GetBranches)
	r.GET("/api/partyrooms/:branch_id", controllers.GetPartyrooms)
	r.GET("/api/selectedroom/:room_id", controllers.GetSelectedRoom)
	r.POST("/api/select-time", controllers.SaveSelectTime)
	r.GET("/api/foodoptions", controllers.Getfoodoptions)
	r.GET("/api/bookings/check", controllers.CheckBookingAvailability)
	r.POST("/api/save-party/:branch_id", controllers.SaveParty)
	//파티 정보 불러오기
	r.GET("/api/get-party", controllers.GetParty)
	r.GET("/api/menu/:branch_id", controllers.GetMenu)
	r.GET("/api/price/:branch_id", controllers.GetPrice)
	r.GET("/api/faq", controllers.GetFAQ)
	r.POST("/api/addmenu/:branch_id", controllers.AddMenu)

	//파티 컨펌
	r.POST("/api/confirm-party", models.ConfirmPartyByID)
	// 엔드포인트 확인
	/*
		r.GET("/api/branches", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "왜 안되는 거야 Branch data fetched successfully!",
			})
		})
	*/

	// CORS 미들웨어 설정
	//r.Use(cors.Default()) // 기본 설정 (모든 도메인 허용)

	//r.Run(":8081") // 백엔드 서버 실행
	port := os.Getenv("PORT") // ✅ Render에서 자동으로 지정됨
	if port == "" {
		port = "8081" // ✅ 로컬 개발 시 기본 포트
	}

	r.SetTrustedProxies(nil) // 프록시를 신뢰하지 않도록 설정

	fmt.Println("🚀 Listening on port", port)
	r.Run(":" + port) // ✅ 포트를 문자열로 이어 붙여 사용

}
