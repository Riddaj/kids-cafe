package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/models"
)

// 특정 branch_id에 대한 필터 추가
func Getfoodoptions(c *gin.Context) {
	// GetPartyrooms 함수 호출, branchID는 URL 파라미터로 받음
	foodoptions, err := models.Getfoodoptions(c)
	if err != nil {
		// GetPartyrooms에서 error가 발생한 경우
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 데이터가 없으면 404 Not Found 응답
	if foodoptions == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No foodoptions found for this branch"})
		return
	}

	// 정상적으로 partyrooms 데이터를 반환
	c.JSON(http.StatusOK, gin.H{"foodoptions": foodoptions})

}

// 푸드 옵션을 보여주는 API
/* func Getfoodoptions(c *gin.Context) {
	var foodoptions []models.foodoption
	result := config.DB.Find(&foodoptions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	// 데이터가 정상적으로 가져와지면 결과 출력
	fmt.Println("foodoptions fetched successfully:", foodoptions)

	c.JSON(http.StatusOK, foodoptions)
}
*/
