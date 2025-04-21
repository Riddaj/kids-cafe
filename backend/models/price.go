package models


import (
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)


type Price struct {
	BranchID       string   `gorm:"type:string;" firestore:"branch_id"`
	Duration     string  	 `gorm:"type:string;" firestore:"duration"`
	PriceID string   		`gorm:"type:string;" firestore:"price_id"`
	WeekdayPrice  string 	`gorm:"type:string;" firestore:"weekday_price"`
	WeekendPrice  string   `gorm:"type:string;" firestore:"weekend_price"`
	
}

// 테이블 이름을 명시적으로 설정
func (Price) TableName() string {
	return "Price" // 실제 테이블 이름
}

// GetPrice 함수: Firestore에서 전체 메뉴 목록 가져오기
func GetPrice(ctx *gin.Context, client *firestore.Client) ([]Price, error) {
	var prices []Price
	branchID := ctx.Param("branch_id")
	log.Printf("############ branchID =%s", branchID)

	// price 컬렉션 가져오기
	docs, err := client.Collection("price").
    Where("branch_id", "==", branchID).
    Documents(ctx.Request.Context()).
    GetAll();

	if err != nil {
		return nil, fmt.Errorf("failed to get menu documents: %v", err)
	}

	// 각 문서를 Menu 구조체로 변환
	for _, doc := range docs {
		var price Price
		if err := doc.DataTo(&price); err != nil {
			log.Printf("⚠️ 메뉴 파싱 실패: %v", err)
			continue
		}
		prices = append(prices, price)
	}

	return prices, nil
}
