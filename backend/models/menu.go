package models

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	BranchID     string   `gorm:"type:int;" firestore:"branch_id"`
	Price		 float64	`gorm:type:int;" firestore:"price"`
	Flavors      []string `firestore:"flavors"`
}

// 테이블 이름을 명시적으로 설정
func (Menu) TableName() string {
	return "Menu" // 실제 테이블 이름
}

// GetMenu 함수: Firestore에서 전체 메뉴 목록 가져오기
func GetMenu(ctx *gin.Context, client *firestore.Client) ([]Menu, error) {
	var menus []Menu
	branchID := ctx.Param("branch_id")
	log.Printf("############ branchID =%s", branchID)

	// menu 컬렉션 가져오기
	docs, err := client.Collection("menu").
	Where("branch_id", "==", branchID).
	Documents(ctx.Request.Context()).
	GetAll();
	
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

// AddMenu 함수: Firestore에서 전체 메뉴 목록 가져오기
func AddMenu(ctx *gin.Context, client *firestore.Client) (Menu, error) {

	var menu Menu

		// JSON 데이터 바인딩
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "요청 파싱 실패"})
		return menu, err
	}

	branchID := ctx.Param("branch_id")
	log.Printf("############ branchID =%s", branchID)
	menu.BranchID = branchID
	
	// MenuID 자동 생성 (예: UUID 등)
	menu.MenuID = uuid.NewString()
	log.Printf("@#######등록할 메뉴 ::: %+v", menu)

	_, err := client.Collection("menu").Doc(menu.MenuID).Set(ctx, menu)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"error": "등록 실패"})
		return menu, err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "메뉴 등록 완료",
		"menu_id": menu.MenuID,
	})

	return menu, nil
}
