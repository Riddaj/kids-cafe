package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	// DB 연결을 가져오기 위함
	// Branch 모델을 가져오기 위함
	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/config"
	"github.com/johnnydev/kids-cafe-backend/models"
)

// 모든 지점(branch) 목록을 가져오는 API (특정 branch_id에 대한 필터 추가)
func GetPartyrooms(c *gin.Context) {
	// URL 파라미터에서 branch_id 가져오기
	branchID := c.Param("branch_id")

	// branch_id 값이 없으면 400 Bad Request 응답
	if branchID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_id is required"})
		return
	}

	// branchID를 숫자로 변환
	id, err := strconv.Atoi(branchID) // branchID를 int로 변환
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch_id"})
		return
	}

	var partyrooms []models.Partyroom
	result := config.DB.Where("branch_id = ?", id).Find(&partyrooms)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 데이터가 정상적으로 가져와지면 결과 출력
	fmt.Println("&&&&& let's go party now work that body partyrooms fetched successfully &&&&& :", partyrooms)

	c.JSON(http.StatusOK, partyrooms)

}
