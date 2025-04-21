package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/models"
)

// 특정 branch_id에 대한 필터 추가
func GetFAQ(c *gin.Context) {
	// GetPartyrooms 함수 호출, branchID는 URL 파라미터로 받음
	FAQs, err := models.GetFAQ(c)
	if err != nil {
		// GetPartyrooms에서 error가 발생한 경우
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 데이터가 없으면 404 Not Found 응답
	if FAQs == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No FAQs found for this FAQs"})
		return
	}

	// 정상적으로 partyrooms 데이터를 반환
	c.JSON(http.StatusOK, gin.H{"FAQs": FAQs})

}
