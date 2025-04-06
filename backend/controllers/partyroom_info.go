package controllers

import (
	// DB 연결을 가져오기 위함
	// Branch 모델을 가져오기 위함

	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
	"github.com/johnnydev/kids-cafe-backend/models"
	"google.golang.org/api/iterator"
)

// 모든 지점(branch) 목록을 가져오는 API (특정 branch_id에 대한 필터 추가)
func GetPartyrooms(c *gin.Context) {
	// GetPartyrooms 함수 호출, branchID는 URL 파라미터로 받음
	partyrooms, err := models.GetPartyrooms(c)
	if err != nil {
		// GetPartyrooms에서 error가 발생한 경우
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 데이터가 없으면 404 Not Found 응답
	if partyrooms == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No partyrooms found for this branch"})
		return
	}

	// 정상적으로 partyrooms 데이터를 반환
	c.JSON(http.StatusOK, gin.H{"partyrooms": partyrooms})

}

// ################# 선택된 파티룸만 가져오는 함수 #################
// ################# 선택된 파티룸만 가져오는 함수 #################

func GetSelectedRoom(c *gin.Context) {

	// URL 경로에서 room_id 추출
	roomID := c.Param("room_id")
	if roomID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room_id가 제공되지 않았습니다."})
		return
	}

	// Firestore 클라이언트 가져오기
	client, err := firebase.GetFirestoreClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Firestore 클라이언트를 초기화할 수 없습니다."})
		return
	}

	// Firestore에서 roomID에 해당하는 문서 조회
	iter := client.Collection("partyrooms").Where("room_id", "==", roomID).Documents(context.Background())
	defer iter.Stop()

	var partyroom models.Partyroom
	found := false
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "파티룸 정보를 가져오는 데 실패했습니다."})
			return
		}

		if err := doc.DataTo(&partyroom); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "파티룸 데이터 파싱에 실패했습니다."})
			return
		}
		found = true
		break // 첫 번째 매칭되는 문서만 가져옵니다
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"message": "해당 room_id를 가진 파티룸을 찾을 수 없습니다."})
		return
	}

	// 파티룸 정보 반환
	c.JSON(http.StatusOK, gin.H{"partyroom": partyroom})

}

/*
roomID := c.Param("room_id")                // URL 경로에서 ID 가져오기
	branchID := c.DefaultQuery("branch_id", "") // Query String에서 가져오기
	roomName := c.DefaultQuery("room_name", "") // Query String에서 가져오기

	fmt.Println("@@@ branchID:", branchID, "@@@ roomName:", roomName, "@@@ roomID:", roomID) // 디버깅용

	if branchID == "" || roomName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing parameters"})
		return
	}

	partyrooms, err := models.GetPartyrooms(context.Background(), branchID, roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(partyrooms) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No party room found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"partyroom": partyrooms[0]})

*/

// 특정 branch_id에 해당하는 partyrooms만 필터링
/*

	var filteredPartyrooms []models.Partyroom
	for _, partyroom := range allPartyrooms {
		if partyroom.BranchID == branchID {
			filteredPartyrooms = append(filteredPartyrooms, partyroom)
		}
	}

	// 결과 응답 반환
	if len(filteredPartyrooms) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No partyrooms found for this branch"})
		} else {
			c.JSON(http.StatusOK, gin.H{"partyrooms": filteredPartyrooms})
		}
*/

/*
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
	//fmt.Println("###왜 안나오냐고#### branchID:", branchID)
	//fmt.Println("###왜 안나오냐고#### roomName:", roomName)
	//fmt.Println("###왜 안나오냐고#### roomId:", roomId)

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
*/
