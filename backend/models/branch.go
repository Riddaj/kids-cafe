package models

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
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

func GetBranches(db *gorm.DB) ([]Branch, error) {
	var branches []Branch
	err := db.Find(&branches).Error

	fmt.Println("#### 데이터 나와라나와라 얍 #### Fetched branches:", branches)

	if err != nil {
		fmt.Println("#### Error fetching branches:", err)
	} else {
		fmt.Println("#### 아 제발요 #### Fetched branches:", branches)
	}

	// ✅ 데이터가 비어있는지 확인
	if len(branches) == 0 {
		fmt.Println("⚠️ No branches found in database")
	} else {
		fmt.Println("✅ Fetched branches:", branches) // 🔥 디버깅 출력문
	}

	return branches, err
}
