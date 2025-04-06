package models

type Party struct {
	BranchID        string   `gorm:"type:string;" firestore:"branch_id"`
	PartyID         string   `gorm:"type:string" firestore:"party_id"`
	PartyroomID     string   `gorm:"type:string" firestore:"partyroom_id"`
	PartyroomName   string   `gorm:"type:string" firestore:"partyroom_name"`
	Partytime       string   `gorm:"type:string" firestore:"partytime"`
	UserID          string   `gorm:"type:string" firestore:"user_id"`
	Weekday         string   `gorm:"type:string" firestore:"weekday"`
	TotalPrice      int      `gorm:"type:int" firestore:"total_price"`
	PartyOwnerPhone int      `gorm:"type:int" firestore:"owner_phone"`
	SelectedFood    []string `gorm:"type:json" firestore:"selected_food"`
	PartyKidName    string   `gorm:"type:string" firestore:"kid_name"`
	PartyOwnerName  string   `gorm:"type:string" firestore:"owner_name"`
	Branch          Branch   `firestore:"branches"`
}

// 테이블 이름을 명시적으로 설정
func (Party) TableName() string {
	return "Party" // 실제 테이블 이름
}
