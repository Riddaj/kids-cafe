package models

import (
	"fmt"

	"gorm.io/gorm"
)

// Branch 모델 정의
type Partyroom struct {
	ID                  string `gorm:"primaryKey;type:varchar(20)" json:"id"`
	RoomName            string `gorm:"type:varchar(100);not null" json:"room_name"`
	RoomDeposit         int    `gorm:"type:int" json:"room_deposit"`
	RoomPriceWeekday    int    `gorm:"type:int" json:"room_price_weekday"`
	RoomPriceWeekend    int    `gorm:"type:int" json:"room_price_weekend"`
	Capacity            int    `gorm:"type:int" json:"capacity"`
	AdditionalHeadcount int    `gorm:"type:int" json:"additional_headcount"`
	BranchID            int    `gorm:"type:int;not null" json:"branch_id"`
	Branch              Branch `json:"branch"` // 관계 설정

	// description 필드 추가
	Description string `gorm:"type:text" json:"description"`

	// 관계 설정 (foreign key)
	//Branch Branch `gorm:"foreignKey:BranchID;references:ID" json:"branch"`
}

func (Partyroom) TableName() string {
	return "partyrooms" // ✅ GORM이 올바른 테이블을 찾도록 지정
}

func GetPartyrooms(db *gorm.DB, branchID int) ([]Partyroom, error) {
	var partyrooms []Partyroom
	//err := db.Find(&partyrooms).Error
	//err := db.Debug().Preload("Branch").Find(&partyrooms).Error
	err := db.Debug().Preload("Branch").Where("branch_id = ?", branchID).Find(&partyrooms).Error

	//250217
	//err := db.Select("partyrooms.*, branches.id as branch_id, branches.branch_name, branches.email, branches.phone, branches.location, branches.bank_bsb, branches.bank_account").
	//	Joins("LEFT JOIN branches ON branches.id = partyrooms.branch_id").
	//	Find(&partyrooms).Error

	//err := db.Joins("LEFT JOIN branches ON branches.id = partyrooms.branch_id").Find(&partyrooms).Error

	fmt.Println("#### let's go party #### Fetched partyrooms:", partyrooms)

	if err != nil {
		fmt.Println("#### Error fetching partyrooms:", err)
	} else {
		fmt.Println("#### let's go party #### Fetched partyrooms:", partyrooms)
	}

	// ✅ 데이터가 비어있는지 확인
	if len(partyrooms) == 0 {
		fmt.Println("⚠️ No branches found in database")
	} else {
		fmt.Println("✅ Fetched branches:", partyrooms) // 🔥 디버깅 출력문
	}

	return partyrooms, err
}
