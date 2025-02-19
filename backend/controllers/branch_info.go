package controllers

import (
	"fmt"
	"net/http"

	// DB 연결을 가져오기 위함
	// Branch 모델을 가져오기 위함
	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/config"
	"github.com/johnnydev/kids-cafe-backend/models"
)

// 모든 지점(branch) 목록을 가져오는 API
func GetBranches(c *gin.Context) {
	var branches []models.Branch
	result := config.DB.Find(&branches)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 데이터가 정상적으로 가져와지면 결과 출력
	fmt.Println("Branches fetched successfully:", branches)

	c.JSON(http.StatusOK, branches)
}
