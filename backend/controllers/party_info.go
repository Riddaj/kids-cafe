package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
	"github.com/johnnydev/kids-cafe-backend/models"
)

func CheckBookingAvailability(c *gin.Context) {
	// 클라이언트에서 받은 날짜와 시간 파라미터
	date := c.DefaultQuery("date", "")
	time := c.DefaultQuery("time", "")

	// 파라미터가 없으면 에러 처리
	if date == "" || time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "날짜와 시간이 필요합니다."})
		return
	}

	// Firebase 초기화
	firebase.InitializeFirebase()

	// Firestore 클라이언트 가져오기
	client, err := firebase.GetFirestoreClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("🎇🎈🎉Error connecting to Firestore: %v", err)})
		return
	}

	// Firestore에서 해당 날짜와 시간에 예약된 문서가 있는지 확인
	query := client.Collection("party").Where("partydate", "==", date).Where("partytime", "==", time)
	docs, err := query.Documents(context.Background()).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Firestore 쿼리 오류: %v", err)})
		return
	}

	// 예약이 없으면 true, 있으면 false 반환
	if len(docs) > 0 {
		c.JSON(http.StatusOK, gin.H{"available": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"available": true})
	}
}

// 파티 예약 insert
func SaveParty(c *gin.Context) {
	fmt.Print("✅✅ save party 확인 ✅✅ ")

	// Firestore 클라이언트 가져오기
	client, err := firebase.GetFirestoreClient() // firebase.GetFirestoreClient는 Firestore 클라이언트를 리턴하는 함수
	if err != nil {
		fmt.Println("❌ Firestore 클라이언트 생성 실패:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	_, err = models.SaveParty(c, client)
	//_, err = models.AddMenu(c, client)

	if err != nil {
		// GetPartyrooms에서 error가 발생한 경우
		fmt.Println("❌ SaveParty 실패:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 메뉴 추가 성공시, 응답 메시지만 보내기
	//c.JSON(http.StatusOK, gin.H{
	//	"success": true,
	//	"message": "파티가 성공적으로 추가되었습니다.",
	//})

}

// 모든 지점(branch) 목록을 가져오는 API (특정 branch_id에 대한 필터 추가)
func GetParty(c *gin.Context) {
	// GetPartyrooms 함수 호출, branchID는 URL 파라미터로 받음
	parties, err := models.GetParty(c)
	if err != nil {
		// GetPartyrooms에서 error가 발생한 경우
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 데이터가 없으면 404 Not Found 응답
	if parties == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No parties found for this branch"})
		return
	}

	// 정상적으로 partyrooms 데이터를 반환
	c.JSON(http.StatusOK, gin.H{"parties": parties})

}

// if c.Request.Method == http.MethodPost {
// 	var party models.Party

// 	// JSON 요청 바디 파싱
// 	if err := json.NewDecoder(c.Request.Body).Decode(&party); err != nil {
// 		fmt.Println("❌ JSON 파싱 에러:", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	fmt.Printf("📌 받은 데이터: %+v\n", party) // 🔴 디버깅 로그 추가

// 	// Firebase 초기화
// 	firebase.InitializeFirebase()

// 	// Firestore 클라이언트 가져오기
// 	client, err := firebase.GetFirestoreClient()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error connecting to Firestore: %v", err)})
// 		return
// 	}

// 	// Firestore에 파티 정보 저장
// 	_, _, err = client.Collection("party").Add(context.Background(), party)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error saving party to Firestore: %v", err)})
// 		return
// 	}

// 	// 성공적인 응답
// 	c.JSON(http.StatusOK, gin.H{"message": "Party booking saved successfully"})
// } else {
// 	c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
// }
