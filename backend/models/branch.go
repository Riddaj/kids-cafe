package models

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
)

// Branch 모델 정의
type Branch struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Branch_name  string         `json:"branch_name"`
	Email        string         `json:"email"`
	Phone        string         `json:"phone"`
	Location     string         `json:"location"`
	Insta        sql.NullString `json:"sns_insta"`
	Facebook     sql.NullString `json:"sns_facebook"`
	Bank_bsb     int            `json:"bank_bsb"`
	Bank_account int            `json:"bank_account"`

	// Partyroom 관계 설정
	Partyrooms []Partyroom `gorm:"foreignKey:BranchID;references:ID" json:"partyrooms"`
}

func (Branch) TableName() string {
	return "branches" // ✅ GORM이 올바른 테이블을 찾도록 지정
}

// Firebase에서 branches 데이터를 가져오는 함수
func GetBranches(c *gin.Context) {
	client, err := firebase.GetFirestoreClient() // Firestore 클라이언트 가져오기
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Firestore client"})
		return
	}
	defer client.Close() // 클라이언트 종료

	// Firestore에서 "branches" 컬렉션의 모든 문서 가져오기
	iter := client.Collection("branches").Documents(c)
	var branches []Branch

	// Firestore 문서 반복문
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		// Firestore 문서 데이터를 Branch 구조체로 변환
		var branch Branch
		doc.DataTo(&branch)

		// Firestore의 ID는 string이므로 uint로 변환
		idUint, err := strconv.ParseUint(doc.Ref.ID, 10, 32) // 10진수, 32비트 변환
		if err != nil {
			log.Println("Error converting Firestore ID to uint:", err)
		} else {
			branch.ID = uint(idUint)
		}
		//branch.ID = doc.Ref.ID                               // 문서의 ID를 Branch 구조체에 추가

		branches = append(branches, branch) // branches 슬라이스에 추가
	}

	// 데이터 확인
	if len(branches) == 0 {
		fmt.Println("⚠️ No branches found in Firestore")
		c.JSON(http.StatusOK, gin.H{"message": "No branches found"})
		return
	} else {
		fmt.Println("✅ ###### Fetched branches: ######## : ", branches)
		c.JSON(http.StatusOK, gin.H{"branches": branches}) // 데이터 반환
	}
}

//250310 이전
// func GetBranches(db *gorm.DB) ([]Branch, error) {
// 	var branches []Branch
// 	err := db.Find(&branches).Error

// 	fmt.Println("#### 데이터 나와라나와라 #### Fetched branches:", branches)

// 	if err != nil {
// 		fmt.Println("#### Error fetching branches:", err)
// 	} else {
// 		fmt.Println("#### 아 제발요 #### Fetched branches:", branches)
// 	}

// 	// ✅ 데이터가 비어있는지 확인
// 	if len(branches) == 0 {
// 		fmt.Println("⚠️ No branches found in database")
// 	} else {
// 		fmt.Println("✅ Fetched branches:", branches) // 🔥 디버깅 출력문
// 	}

// 	return branches, err
// }
