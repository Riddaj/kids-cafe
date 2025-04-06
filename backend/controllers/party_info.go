package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
	"github.com/johnnydev/kids-cafe-backend/models"
)

func SaveParty(c *gin.Context) {
	fmt.Print("✅✅ save party 확인 ✅✅ ")
	if c.Request.Method == http.MethodPost {
		var party models.Party

		// JSON 요청 바디 파싱
		if err := json.NewDecoder(c.Request.Body).Decode(&party); err != nil {
			fmt.Println("❌ JSON 파싱 에러:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("📌 받은 데이터: %+v\n", party) // 🔴 디버깅 로그 추가

		// Firebase 초기화
		firebase.InitializeFirebase()

		// Firestore 클라이언트 가져오기
		client, err := firebase.GetFirestoreClient()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error connecting to Firestore: %v", err)})
			return
		}

		// Firestore에 파티 정보 저장
		_, _, err = client.Collection("party").Add(context.Background(), party)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving party to Firestore: %v", err)})
			return
		}

		// 성공적인 응답
		c.JSON(http.StatusOK, gin.H{"message": "Party booking saved successfully"})
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
	}
}
