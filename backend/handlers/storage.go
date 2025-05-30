package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/johnnydev/kids-cafe-backend/firebase"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	fmt.Println("📥 UploadHandler 진입 확인") // ✅ 가장 처음
	// 1. 파일 읽기
	file, header, err := r.FormFile("depositImage")
	if err != nil {
		fmt.Println("❌ 파일 읽기 실패:", err) // ✅ 에러 내용 확인
		http.Error(w, "file read error", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 2. Firebase Storage 연결
	client, err := firebase.GetApp().Storage(ctx)
	if err != nil {
		http.Error(w, "storage init error", http.StatusInternalServerError)
		return
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		fmt.Println("❌ Bucket 가져오기 실패:", err)
		http.Error(w, "bucket error", http.StatusInternalServerError)
		return
	}

	// 3. 파일 업로드
	filename := fmt.Sprintf("deposits/%d_%s", time.Now().Unix(), header.Filename)
	fmt.Println("🚀 업로드할 파일명:", filename)
	wc := bucket.Object(filename).NewWriter(ctx)
	wc.ContentType = header.Header.Get("Content-Type")

	if _, err := io.Copy(wc, file); err != nil {
		fmt.Println("❌ 파일 업로드 실패:", err)
		http.Error(w, "upload error", http.StatusInternalServerError)
		return
	}

	if err := wc.Close(); err != nil {
		fmt.Println("❌ 파일 저장 마무리 실패:", err)
		http.Error(w, "upload finalize error", http.StatusInternalServerError)
		return
	}
	fmt.Println("✅ 파일 업로드 성공")

	rawJSON := r.FormValue("bookingData")
	fmt.Println("📦 raw bookingData string:", rawJSON)

	// 4. 예약 정보 JSON 읽기
	var bookingData map[string]interface{}
	if err := json.Unmarshal([]byte(r.FormValue("bookingData")), &bookingData); err != nil {
		fmt.Println("❌ JSON 파싱 실패:", err)
		http.Error(w, "booking data error", http.StatusBadRequest)
		return
	}

	fmt.Println("✅ 예약 정보 JSON:", bookingData)

	// 5. 이미지 URL 만들기
	bucketName := "kids-cafe-booking-project.firebasestorage.app" // 🔁 실제 버킷 이름으로 바꿔주세요!
	imageURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media",
		bucketName, url.PathEscape(filename))

	fmt.Println("✅ 생성된 이미지 URL:", imageURL)

	// 6. 예약 정보에 이미지 URL 추가
	bookingData["party_id"] = uuid.NewString() // ✅ party_id 생성
	bookingData["deposit_image_url"] = imageURL
	bookingData["is_confirmed"] = false // 초기값 설정

	// 7. Firestore에 저장
	firestoreClient, err := firebase.GetFirestoreClient()
	if err != nil {
		fmt.Println("❌ Firestore 클라이언트 연결 실패:", err)
		http.Error(w, "firestore init error", http.StatusInternalServerError)
		return
	}
	defer firestoreClient.Close()

	_, _, err = firestoreClient.Collection("party").Add(ctx, bookingData)
	if err != nil {
		http.Error(w, "firestore save error", http.StatusInternalServerError)
		return
	}
	fmt.Println("🎉 Firestore 저장 성공!")
	// 8. 응답
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Uploaded","imagePath":"` + imageURL + `"}`))
}
