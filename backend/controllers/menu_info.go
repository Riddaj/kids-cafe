package controllers

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
	"github.com/johnnydev/kids-cafe-backend/models"
)

// 특정 branch_id에 대한 필터 추가
func GetMenu(c *gin.Context) {
	log.Printf("############ GetMenu 호출")

	// Firestore 클라이언트 가져오기
	client, err := firebase.GetFirestoreClient() // firebase.GetFirestoreClient는 Firestore 클라이언트를 리턴하는 함수
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	// GetPartyrooms 함수 호출, branchID는 URL 파라미터로 받음
	menus, err := models.GetMenu(c, client)

	if err != nil {
		// GetPartyrooms에서 error가 발생한 경우
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 데이터가 없으면 404 Not Found 응답
	if menus == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No menus found for this branch"})
		return
	}

	// 정상적으로 partyrooms 데이터를 반환
	c.JSON(http.StatusOK, gin.H{"menus": menus})

}

//메뉴 추가
func AddMenu(c *gin.Context) {
	log.Printf("############ AddMenu 호출")

	// Firestore 클라이언트 가져오기
	client, err := firebase.GetFirestoreClient() // firebase.GetFirestoreClient는 Firestore 클라이언트를 리턴하는 함수
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer client.Close()

	_, err = models.AddMenu(c, client)

	if err != nil {
		// GetPartyrooms에서 error가 발생한 경우
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 메뉴 추가 성공시, 응답 메시지만 보내기
	c.JSON(http.StatusOK, gin.H{
		"message": "메뉴가 성공적으로 추가되었습니다.",
	})

}
