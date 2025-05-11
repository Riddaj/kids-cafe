// firebase/firebase.go

package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"cloud.google.com/go/firestore" // firestore.Client 임포트
)

var app *firebase.App

// Firebase 초기화
func InitializeFirebase() {
	serviceAccount := "/etc/secrets/kids-cafe-booking-project-firebase-adminsdk-fbsvc-e1544abf3a.json"
	opt := option.WithCredentialsFile(serviceAccount)
	var err error
	app, err = firebase.NewApp(context.Background(), nil, opt)
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
