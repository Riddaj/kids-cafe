package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/config"
	"github.com/johnnydev/kids-cafe-backend/models"
)

// 푸드 옵션을 보여주는 API
func GetFoodOptions(c *gin.Context) {
	var foodOptions []models.FoodOption
	result := config.DB.Find(&foodOptions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	// 데이터가 정상적으로 가져와지면 결과 출력
	fmt.Println("foodOptions fetched successfully:", foodOptions)

	c.JSON(http.StatusOK, foodOptions)
}
