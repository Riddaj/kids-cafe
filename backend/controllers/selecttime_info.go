package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/config"
	"github.com/johnnydev/kids-cafe-backend/models"
)

// 날짜 데이터를 받아서 DB에 저장하는 API
func SaveSelectTime(c *gin.Context) {
	var request struct {
		PartyroomID  string `json:"partyroom_id"`
		SelectedDate string `json:"selected_date"` // 받아오는 날짜 및 시간
		Selectedtime string `json:"selected_time"`
	}

	// 요청 데이터를 바인딩
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// 받은 데이터 출력 (디버깅용)
	log.Println("Received PartyroomID:", request.PartyroomID)
	log.Println("Received Date:", request.SelectedDate)
	log.Println("Received Selected Time:", request.Selectedtime)

	//Wed Feb 05 2025 00:00:00

	// 문자열을 공백 기준으로 분리
	dateStr := strings.Split(request.SelectedDate, " ")

	// "Wed" -> 요일
	dayOfWeek := dateStr[0]

	// dateStr[1] (월), dateStr[2] (날짜), dateStr[3] (연도)를 결합하여 "MM dd yyyy" 형식으로 날짜 만들기
	formattedDate := dateStr[1] + "-" + dateStr[2] + "-" + dateStr[3] // 예: "Feb 05 2025"

	// 날짜 포맷 파싱: "Mon Jan 02 2006 15:04:05"
	// parsedTime, err := time.Parse("request.DateTime", dateStr[])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 원하는 형식으로 변환: "dd-MM-yyyy"
	//formattedDate := parsedTime.Format("02-01-2006")

	// 결과 출력
	fmt.Println("@@@@@ Formatted Date @@@@@ :", formattedDate)

	// "2025-02-19" -> time.Time 변환
	selectedDate := formattedDate
	// if err != nil {
	// 	log.Println("Error parsing date:", err)
	// 	c.JSON(400, gin.H{"error": "Invalid date format"})
	// 	return
	// }

	//선택한 시간
	time1 := request.Selectedtime

	// DB에 저장
	selectTime := models.SelectTime{
		PartyRoomID:  request.PartyroomID, // 서버에서 받아오는 값
		SelectedDate: selectedDate,        // time.Time 형식으로 변환된 date
		DayOfWeek:    dayOfWeek,
		Time1:        time1, // time.Time 형식으로 변환된 startTime  대신에 time1,2로 변경
		//Time2:       time2, // time.Time 형식으로 변환된 endTime 대신에 time1,2로 변경
	}

	// 모델을 DB에 저장
	if err := config.DB.Create(&selectTime).Error; err != nil {
		log.Println("Error saving select time:", err)
		c.JSON(500, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(200, gin.H{"message": "Data saved successfully"})
}

// 날짜 데이터를 받아서 DB에 저장하는 API
/*
func SaveSelectTime(c *gin.Context) {
	var request struct {
		DateTime string `json:"date_time"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 문자열을 공백 기준으로 분리
	parts := strings.Split(request.DateTime, " ")
	if len(parts) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	date := parts[0]      // "2025-02-19"
	dayOfWeek := parts[1] // "Wed"
	timeRange := parts[2] // "10:00-12:00"

	// 시간 범위를 "-" 기준으로 분리
	times := strings.Split(timeRange, "-")
	if len(times) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid time format"})
		return
	}

	startTime := times[0] // "10:00"
	endTime := times[1]   // "12:00"

	//var selectTime []models.SelectTime
	// DB에 저장
	selectTime := models.SelectTime{
		ID:          "some_unique_id",
		PartyRoomID: "some_party_room_id",
		Date:        date,
		DayOfWeek:   dayOfWeek,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	if err := config.DB.Create(&selectTime).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data saved successfully"})
}
*/
