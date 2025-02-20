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
	"gorm.io/gorm"
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

func GetSelectedRoom(c *gin.Context) {
	roomId := c.Param("room_id")                                      // URL 경로에서 ID 가져오기
	branchID := c.DefaultQuery("branch_id", "")                       // Query String에서 가져오기
	roomName := c.DefaultQuery("room_name", "")                       // Query String에서 가져오기
	fmt.Println("@@@ branchID:", branchID, "@@@ roomName:", roomName) // 디버깅

	//branchID := c.Request.URL.Query().Get("branch_id")
	//roomName := c.Request.URL.Query().Get("room_name")

	// 디버깅용 출력
	fmt.Println("###왜 안나오냐고#### branchID:", branchID)
	fmt.Println("###왜 안나오냐고#### roomName:", roomName)
	fmt.Println("###왜 안나오냐고#### roomId:", roomId)

	if branchID == "" || roomName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
		return
	}

	var room models.Partyroom
	if err := config.DB.Where("branch_id = ? AND id = ?", branchID, roomId).First(&room).Error; err != nil {

		fmt.Println("###  ###   ### Running query with branchID:", branchID, "and roomId:", roomId)

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"********* error *********": "Party room not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, room)
}
