package models

import (
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/johnnydev/kids-cafe-backend/firebase"
)

type Party struct {
	BranchID                  string   `firestore:"branch_id"`
	PartyID                   string   `firestore:"party_id"`
	PartyroomID               string   `json:"partyroom_id" firestore:"partyroom_id"`
	PartyroomName             string   `json:"partyroom_name" firestore:"partyroom_name"`
	Partydate                 string   `firestore:"partydate"`
	Partytime                 string   `firestore:"partytime"`
	PartyroomPrice            int      `json:"partyroom_price" firestore:"partyroom_price"`
	FoodPrice                 int      `json:"food_price" firestore:"food_price"`
	OwnerPhone                string   `json:"owner_phone" firestore:"owner_phone"`
	SelectedFood              []string `json:"selected_food" firestore:"selected_food"`
	KidName                   string   `json:"kid_name" firestore:"kid_name"`
	OwnerName                 string   `json:"owner_name" firestore:"owner_name"`
	KidGender                 string   `json:"kid_gender" firestore:"kid_gender"`
	KidAge                    int      `json:"kid_age" firestore:"kid_age"`
	KidRelation               string   `json:"kid_relation" firestore:"kid_relation"`
	Email                     string   `json:"email" firestore:"email"`
	BalloonDecorationsChecked bool     `json:"balloonDecorationsChecked"`
	BalloonDecorationsTheme   string   `json:"balloonDecorationsTheme"`
	SpecialRequired           []string `json:"special_required" firestore:"special_required"`
	OptionService             string   `json:"option_service" firestore:"option_service"`
	AddRequirement            string   `json:"addRequirement"`
	PaymentMethod             string   `json:"payment_method" firestore:"payment_method"`
	//DepositImageURL           string   `json:"deposit_image_url" firestore:"deposit_image_url"`
	AgreeTerms bool `json:"agree_terms"`
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

	fmt.Println("받은 아이 이름:", party.KidName)
	fmt.Println("저나번호:", party.OwnerPhone)

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

// 모델 함수 - Firestore에서 party 데이터를 가져오는 함수
func GetParty(c *gin.Context) ([]Party, error) {
	//branchIDStr := c.Param("branch_id")
	//branchID, err := strconv.Atoi(branchIDStr) //문자열 int로 변환

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid branch_id"})
	// 	return nil, err // 오류 처리 후 적절히 반환
	// }

	branchID := c.Param("branch_id") // 문자열 그대로 사용

	client, err := firebase.GetFirestoreClient()
	if err != nil {
		// Firestore 클라이언트 연결 실패 시 에러 반환
		return nil, fmt.Errorf("failed to get Firestore client: %v", err)
	}
	defer client.Close() // 클라이언트 종료

	// "partyrooms" 컬렉션의 모든 문서 가져오기
	iter := client.Collection("party").Where("branch_id", "==", branchID).Documents(c)
	var parties []Party

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var party Party
		//doc.DataTo(&partyroom)
		if err := doc.DataTo(&party); err != nil {
			// DataTo에서 오류가 나면 오류 메시지를 출력하고, 해당 문서 건너뛰기
			log.Printf("#### Error while converting document data ####: %v", err)
			continue
		}
		//partyroom.RoomID = doc.Ref.ID // Firestore 문서 ID 저장
		party.PartyID = doc.Data()["party_id"].(string) // 문서 내 room_id 필드 값 가져오기

		parties = append(parties, party)
	}

	// 데이터가 없을 경우
	if len(parties) == 0 {
		return nil, nil // 데이터가 없으면 nil 반환
	}

	// 데이터가 있을 경우
	return parties, nil // 정상적으로 partyrooms 반환
}
