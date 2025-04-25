package models

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
)

// Branch 모델 정의
type FAQ struct {
	FaqID       string `gorm:"type:varchar(10);" firestore:"faq_id"`
	FaqQuestion string `gorm:"type:varchar(100);" firestore:"faq_question"`
	FaqAnswer   string `gorm:"type:varchar(100);" firestore:"faq_answer"`
}

func (FAQ) TableName() string {
	return "FAQ" // ✅ GORM이 올바른 테이블을 찾도록 지정
}

// Firebase에서 branches 데이터를 가져오는 함수
func GetFAQ(c *gin.Context) ([]FAQ, error) {

	client, err := firebase.GetFirestoreClient()
	if err != nil {
		// Firestore 클라이언트 연결 실패 시 에러 반환
		return nil, fmt.Errorf("failed to get Firestore client: %v", err)
	}
	defer client.Close() // 클라이언트 종료

	// "partyrooms" 컬렉션의 모든 문서 가져오기
	iter := client.Collection("faq").Documents(c)
	var FAQs []FAQ

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var FAQ FAQ
		//doc.DataTo(&partyroom)
		if err := doc.DataTo(&FAQ); err != nil {
			// DataTo에서 오류가 나면 오류 메시지를 출력하고, 해당 문서 건너뛰기
			log.Printf("#### Error while converting document data ####: %v", err)
			continue
		}
		FAQ.FaqID = doc.Data()["faq_id"].(string) // Firestore 문서 ID 저장

		FAQs = append(FAQs, FAQ)
	}

	// 데이터가 없을 경우
	if len(FAQs) == 0 {
		log.Println("### No FAQs found in Firestore ###")
		return nil, nil // 데이터가 없으면 nil 반환
	}

	// 데이터가 있을 경우
	return FAQs, nil // 정상적으로 partyrooms 반환
}
