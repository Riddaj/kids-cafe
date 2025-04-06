package models

import (
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type Option struct {
	Size  string  `firestore:"size"`
	Price float64 `firestore:"price"`
}

type Menu struct {
	MenuID       string   `gorm:"type:string;" firestore:"menu_id"`
	MenuName     string   `gorm:"type:string;" firestore:"menu_name"`
	MenuCategory string   `gorm:"type:string;" firestore:"menu_category"`
	MenuOptions  []Option `firestore:"menu_options"`
	Description  string   `gorm:"type:string;" firestore:"description"`
	ImgUrl       string   `gorm:"type:string;" firestore:"imgUrl"`
	BranchID     int      `gorm:"type:int;" firestore:"branch_id"`
}

// 테이블 이름을 명시적으로 설정
func (Menu) TableName() string {
	return "Menu" // 실제 테이블 이름
}

// GetMenu 함수: Firestore에서 전체 메뉴 목록 가져오기
func GetMenu(ctx *gin.Context, client *firestore.Client) ([]Menu, error) {
	var menus []Menu

	// menu 컬렉션 가져오기
	docs, err := client.Collection("menu").Documents(ctx.Request.Context()).GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get menu documents: %v", err)
	}

	// 각 문서를 Menu 구조체로 변환
	for _, doc := range docs {
		var menu Menu
		if err := doc.DataTo(&menu); err != nil {
			log.Printf("⚠️ 메뉴 파싱 실패: %v", err)
			continue
		}
		menus = append(menus, menu)
	}

	return menus, nil
}
