package models

import (
	"fmt"

	"gorm.io/gorm"
)

type FoodOption struct {
	FoodId          uint   `json:"id" gorm:"primaryKey"`
	FoodName        string `gorm:"type:varchar(50);not null" json:"food_name"`
	FoodQuantity    string `gorm:"type:varchar(20);" json:"food_quantity"`
	FoodPackage     string `gorm:"type:char(1);" json:"food_package"`
	FoodForKids     string `gorm:"type:char(1);" json:"food_forkids"`
	FoodDescription string `gorm:"type:varchar(100);" json:"food_description"`
	FoodPrice       int    `gorm:"type:int;" json:"food_price"`
	BranchID        int    `gorm:"type:varchar(100);" json:"branch_id"`
	Branch          Branch `json:"branch"` // 관계 설정
}

func (FoodOption) TableName() string {
	return "foodOptions" // ✅ GORM이 올바른 테이블을 찾도록 지정
}

func GetFoodOptions(db *gorm.DB) ([]FoodOption, error) {
	var foodOptions []FoodOption

	err := db.Find(&foodOptions).Error

	fmt.Println("#### 푸드 옵션 나와라나와라  #### Fetched branches:", foodOptions)

	// ✅ 데이터가 비어있는지 확인
	if len(foodOptions) == 0 {
		fmt.Println("⚠️ No branches found in database")
	} else {
		fmt.Println("✅ Fetched branches:", foodOptions) // 🔥 디버깅 출력문
	}

	return foodOptions, err
}
