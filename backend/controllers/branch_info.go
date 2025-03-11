package controllers

import (
	"net/http"

	// DB 연결을 가져오기 위함
	// Branch 모델을 가져오기 위함
	"github.com/gin-gonic/gin"
	"github.com/johnnydev/kids-cafe-backend/firebase"
)

// 모든 지점(branch) 목록을 가져오는 API
func GetBranches(c *gin.Context) {
	// Firestore 클라이언트 가져오기
	client, err := firebase.GetFirestoreClient()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get Firestore client"})
		return
	}

	// "branches" 컬렉션에서 데이터 조회
	docs, err := client.Collection("branches").Documents(c.Request.Context()).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch branches"})
		return
	}

	var branches []map[string]interface{}
	for _, doc := range docs {
		branches = append(branches, doc.Data())
	}

	c.JSON(http.StatusOK, gin.H{"branches": branches})

	// 250310 firebase 이전 코드
	// var branches []models.Branch
	// result := config.DB.Find(&branches)

	// if result.Error != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	// 	return
	// }

	// // 데이터가 정상적으로 가져와지면 결과 출력
	// fmt.Println("Branches fetched successfully:", branches)

	// c.JSON(http.StatusOK, branches)
}
