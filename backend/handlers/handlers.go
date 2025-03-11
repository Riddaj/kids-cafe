package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/johnnydev/kids-cafe-backend/firebase"
)

func HandleBookingRequest(w http.ResponseWriter, r *http.Request) {
	// Firebase 초기화
	firebase.InitializeFirebase()

	// Firestore 클라이언트 가져오기
	client, err := firebase.GetFirestoreClient()
	if err != nil {
		http.Error(w, "Firestore 클라이언트 가져오기 실패", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// 예약 정보 가져오기
	// 예: 전화번호로 예약 정보 검색
	phone := r.URL.Query().Get("phone")
	doc, err := client.Collection("bookings").Doc(phone).Get(r.Context())
	if err != nil {
		http.Error(w, "예약 정보 조회 실패", http.StatusNotFound)
		return
	}

	// 결과 출력
	response := map[string]interface{}{
		"phone":   doc.Data()["phone"],
		"details": doc.Data()["details"], // 예약 정보
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
