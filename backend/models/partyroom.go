package models

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
)

// Branch 모델 정의
type Partyroom struct {
	RoomID              string `gorm:"primaryKey;type:string" firestore:"room_id"`
	RoomName            string `gorm:"type:varchar(100);not null" firestore:"room_name"`
	RoomDeposit         int    `gorm:"type:int" firestore:"room_deposit"`
	RoomPriceWeekday    int    `gorm:"type:int" firestore:"room_price_weekday"`
	RoomPriceWeekend    int    `gorm:"type:int" firestore:"room_price_weekend"`
	Capacity            int    `gorm:"type:int" firestore:"capacity"`
	AdditionalHeadcount int    `gorm:"type:int" firestore:"additional_headcount"`
	BranchID            int    `gorm:"type:int;not null" firestore:"branch_id"`
	Branch              Branch `firestore:"branches"`

	// description 필드 추가
	Description string `firestore:"description"`

	// 관계 설정 (foreign key)
	//Branch Branch `gorm:"foreignKey:BranchID;references:ID" json:"branch"`
}

func (Partyroom) TableName() string {
	return "partyrooms" // ✅ GORM이 올바른 테이블을 찾도록 지정
}

// 모델 함수 - Firestore에서 Partyrooms 데이터를 가져오는 함수
func GetPartyrooms(c *gin.Context) ([]Partyroom, error) {
	branchIDStr := c.Param("branch_id")
	branchID, err := strconv.Atoi(branchIDStr) //문자열 int로 변환

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch_id"})
		return nil, err // 오류 처리 후 적절히 반환
	}

	client, err := firebase.GetFirestoreClient()
	if err != nil {
		// Firestore 클라이언트 연결 실패 시 에러 반환
		return nil, fmt.Errorf("failed to get Firestore client: %v", err)
	}
	defer client.Close() // 클라이언트 종료

	// "partyrooms" 컬렉션의 모든 문서 가져오기
	iter := client.Collection("partyrooms").Where("branch_id", "==", branchID).Documents(c)
	var partyrooms []Partyroom

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var partyroom Partyroom
		//doc.DataTo(&partyroom)
		if err := doc.DataTo(&partyroom); err != nil {
			// DataTo에서 오류가 나면 오류 메시지를 출력하고, 해당 문서 건너뛰기
			log.Printf("#### Error while converting document data ####: %v", err)
			continue
		}
		//partyroom.RoomID = doc.Ref.ID // Firestore 문서 ID 저장
		partyroom.RoomID = doc.Data()["room_id"].(string) // 문서 내 room_id 필드 값 가져오기

		partyrooms = append(partyrooms, partyroom)
	}

	// 데이터가 없을 경우
	if len(partyrooms) == 0 {
		return nil, nil // 데이터가 없으면 nil 반환
	}

	// 데이터가 있을 경우
	return partyrooms, nil // 정상적으로 partyrooms 반환
}
