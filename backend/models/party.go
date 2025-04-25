package models

import (
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Party struct {
	BranchID        string   `firestore:"branch_id"`
	PartyID         string   `firestore:"party_id"`
	PartyroomID     string   `firestore:"partyroom_id"`
	PartyroomName   string   `firestore:"partyroom_name"`
	Partydate       string   `firestore:"partydate"`
	Partytime       string   `firestore:"partytime"`
	FoodPrice       int      `firestore:"food_price"`
	PartyOwnerPhone int      `firestore:"owner_phone"`
	SelectedFood    []string `firestore:"selected_food"`
	PartyKidName    string   `firestore:"kid_name"`
	PartyOwnerName  string   `firestore:"owner_name"`
	KidGender       string   `firestore:"kid_gender"`
	KidAge          int      `firestore:"kid_age"`
	KidRelation     string   `firestore:"kid_relation"`
	Email           string   `firestore:"email"`
	SpecialRequired []string `firestore:"special_required"`
	OptionService   string   `firestore:"option_service"`
}

// 테이블 이름을 명시적으로 설정
func (Party) TableName() string {
	return "Party" // 실제 테이블 이름
}

// AddMenu 함수: Firestore에서 전체 메뉴 목록 가져오기
func SaveParty(ctx *gin.Context, client *firestore.Client) (Party, error) {

	var party Party

	// JSON 데이터 바인딩
	if err := ctx.ShouldBindJSON(&party); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "요청 파싱 실패"})
		return party, err
	}

	branchID := ctx.Param("branch_id")
	log.Printf("############ branchID =%s", branchID)
	party.BranchID = branchID

	// MenuID 자동 생성 (예: UUID 등)
	party.PartyID = uuid.NewString()
	log.Printf("@#######등록할 파티  ::: %+v", party)

	_, err := client.Collection("party").Doc(party.PartyID).Set(ctx, party)
	if err != nil {
		log.Printf("❌ Firestore에 데이터 저장 실패: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "등록 실패"})
		return party, err
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "파티 등록 완료",
		"party_id": party.PartyID,
	})

	return party, nil
}
