package models

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
)

type foodoption struct {
	FoodId          string `firestore:"food_id" gorm:"primaryKey"`
	FoodName        string `gorm:"type:varchar(50);not null" firestore:"food_name"`
	FoodQuantity    string `gorm:"type:varchar(20);" firestore:"food_quantity"`
	FoodPackage     bool   `gorm:"type:bool;" firestore:"food_package"`
	FoodForKids     string `gorm:"type:char(1);" firestore:"food_forkids"`
	FoodDescription string `gorm:"type:varchar(100);" firestore:"food_description"`
	FoodPrice       int    `gorm:"type:int;" firestore:"food_price"`
	BranchID        string `gorm:"type:string;" firestore:"branch_id"`
	Branch          Branch `firestore:"branch"` // 관계 설정
}

func (foodoption) TableName() string {
	return "foodoptions" // ✅ GORM이 올바른 테이블을 찾도록 지정
}

func Getfoodoptions(c *gin.Context) ([]foodoption, error) {
	/*
		branchIDStr := c.Param("branch_id")
		branchID, err := strconv.Atoi(branchIDStr) //문자열 int로 변환

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch_id"})
			return nil, err // 오류 처리 후 적절히 반환
		}
	*/

	client, err := firebase.GetFirestoreClient()
	if err != nil {
		// Firestore 클라이언트 연결 실패 시 에러 반환
		return nil, fmt.Errorf("failed to get Firestore client: %v", err)
	}
	defer client.Close() // 클라이언트 종료

	// "partyrooms" 컬렉션의 모든 문서 가져오기
	iter := client.Collection("foodoptions").Documents(c)
	var foodoptions []foodoption

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var foodoption foodoption
		//doc.DataTo(&partyroom)
		if err := doc.DataTo(&foodoption); err != nil {
			// DataTo에서 오류가 나면 오류 메시지를 출력하고, 해당 문서 건너뛰기
			log.Printf("#### Error while converting document data ####: %v", err)
			continue
		}
		foodoption.FoodId = doc.Data()["food_id"].(string) // Firestore 문서 ID 저장

		foodoptions = append(foodoptions, foodoption)
	}

	// 데이터가 없을 경우
	if len(foodoptions) == 0 {
		return nil, nil // 데이터가 없으면 nil 반환
	}

	// 데이터가 있을 경우
	return foodoptions, nil // 정상적으로 partyrooms 반환
}

/* func Getfoodoptions(db *gorm.DB) ([]foodoption, error) {
	var foodoptions []foodoption

	err := db.Find(&foodoptions).Error

	fmt.Println("#### 푸드 옵션 나와라나와라  #### Fetched branches:", foodoptions)

	// ✅ 데이터가 비어있는지 확인
	if len(foodoptions) == 0 {
		fmt.Println("⚠️ No branches found in database")
	} else {
		fmt.Println("✅ Fetched branches:", foodoptions) // 🔥 디버깅 출력문
	}

	return foodoptions, err
}
*/
