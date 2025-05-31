// firebase/firebase.go

package firebase

import (
	"context"
	"encoding/json"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"cloud.google.com/go/firestore" // firestore.Client 임포트
)

var app *firebase.App

// Firebase 초기화
func InitializeFirebase() {
	// 🔑 환경에 따라 경로 설정
	// credPath := "./kids-cafe-booking-project-firebase-adminsdk-fbsvc-e1544abf3a.json" // 기본 로컬용
	// if os.Getenv("ENV") == "production" {
	// 	credPath = "/etc/secrets/kids-cafe-booking-project-firebase-adminsdk-fbsvc-e1544abf3a.json" // Render 배포용
	// }

	credPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")
	if credPath == "" {
		log.Fatal("❌ FIREBASE_CREDENTIALS_PATH 환경변수가 설정되지 않았습니다.")
	}

	// ✅ StorageBucket 설정 추가
	conf := &firebase.Config{
		StorageBucket: "kids-cafe-booking-project.firebasestorage.app",
	}

	//serviceAccount := "/etc/secrets/kids-cafe-booking-project-firebase-adminsdk-fbsvc-e1544abf3a.json"
	opt := option.WithCredentialsFile(credPath)

	var err error
	app, err = firebase.NewApp(context.Background(), conf, opt)
	if err != nil {
		log.Fatalf("Firebase 앱 초기화 실패: %v", err)
	}

}

// Firestore 클라이언트 가져오기
func GetFirestoreClient() (*firestore.Client, error) {
	client, err := app.Firestore(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetApp() *firebase.App {
	return app
}

// 🔐 Firebase 서비스 계정 키 JSON 구조체
type ServiceAccount struct {
	Type        string `json:"type"`
	ProjectID   string `json:"project_id"`
	PrivateKey  string `json:"private_key"`
	ClientEmail string `json:"client_email"`
}

// 🔍 JSON 파일에서 서비스 계정 정보 불러오기
func LoadServiceAccount(path string) (*ServiceAccount, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var sa ServiceAccount
	if err := json.Unmarshal(bytes, &sa); err != nil {
		return nil, err
	}

	return &sa, nil
}
