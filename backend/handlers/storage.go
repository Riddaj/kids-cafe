package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/johnnydev/kids-cafe-backend/firebase"
	"google.golang.org/api/option"
)

func GetSignedURLHandler(c *gin.Context) {
	ctx := context.Background()

	filename := c.Query("filename") // 예: deposits/1748613450_party.png
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "filename required"})
		return
	}

	// 🔐 환경 변수에서 서비스 계정 키 경로 받아오기
	credPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credPath == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "missing FIREBASE_CREDENTIALS_PATH env var"})
		return
	}

	// 🔧 Firebase Storage bucket name
	bucketName := "kids-cafe-booking-project.firebasestorage.app"

	// 🔑 service account 정보로 storage client 생성
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credPath))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "storage client error"})
		return
	}
	defer client.Close()

	// 🔓 JSON 파일에서 GoogleAccessID와 PrivateKey 가져오기
	sa, err := firebase.LoadServiceAccount(credPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load service account"})
		return
	}

	expires := time.Now().Add(1 * time.Hour)
	signedURL, err := storage.SignedURL(bucketName, filename, &storage.SignedURLOptions{
		GoogleAccessID: sa.ClientEmail,
		PrivateKey:     []byte(sa.PrivateKey),
		Method:         "GET",
		Expires:        expires,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate signed URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"signed_url": signedURL,
	})
}

func UploadHandler(c *gin.Context) {

	ctx := context.Background()

	fmt.Println("📥 UploadHandler 진입 확인") // ✅ 가장 처음
	// 1. 파일 읽기
	fileHeader, err := c.FormFile("depositImage")
	if err != nil {
		fmt.Println("❌ 파일 읽기 실패:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "file read error"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("❌ 파일 열기 실패:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "file open error"})
		return
	}
	defer file.Close()

	// 2. Firebase Storage 연결
	client, err := firebase.GetApp().Storage(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "storage init error"})
		return
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		fmt.Println("❌ Bucket 가져오기 실패:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "bucket error"})
		return
	}

	// 3. 파일 업로드
	filename := fmt.Sprintf("deposits/%d_%s", time.Now().Unix(), fileHeader.Filename)

	fmt.Println("🚀 업로드할 파일명:", filename)

	wc := bucket.Object(filename).NewWriter(ctx)
	wc.ContentType = fileHeader.Header.Get("Content-Type")

	if _, err := io.Copy(wc, file); err != nil {
		fmt.Println("❌ 파일 업로드 실패:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upload error"})
		return
	}

	if err := wc.Close(); err != nil {
		fmt.Println("❌ 파일 저장 마무리 실패:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "upload finalize error"})
		return
	}
	fmt.Println("✅ 파일 업로드 성공")

	// 4. JSON 읽기
	rawJSON := c.PostForm("bookingData")
	fmt.Println("📦 raw bookingData string:", rawJSON)

	var bookingData map[string]interface{}
	if err := json.Unmarshal([]byte(rawJSON), &bookingData); err != nil {
		fmt.Println("❌ JSON 파싱 실패:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "booking data error"})
		return
	}

	// 5. 이미지 URL 만들기
	//bucketName := "kids-cafe-booking-project.firebasestorage.app" // ✅ 실제 Firebase 콘솔에서 확인한 주소로 수정
	//imageURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media",
	//	bucketName, url.PathEscape(filename))

	bookingData["party_id"] = uuid.NewString()
	bookingData["deposit_filename"] = filename
	//bookingData["deposit_image_url"] = imageURL
	bookingData["is_confirmed"] = false

	// 6. Firestore 저장
	firestoreClient, err := firebase.GetFirestoreClient()
	if err != nil {
		fmt.Println("❌ Firestore 클라이언트 연결 실패:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "firestore init error"})
		return
	}
	defer firestoreClient.Close()

	_, _, err = firestoreClient.Collection("party").Add(ctx, bookingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "firestore save error"})
		return
	}

	fmt.Println("🎉 Firestore 저장 성공!")
	c.JSON(http.StatusOK, gin.H{
		"message":  "Uploaded",
		"filename": filename,
	})
}
