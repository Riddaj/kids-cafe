package models

type SelectTime struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	PartyRoomID  string    `gorm:"column:partyroom_id" json:"partyroom_id"`
	SelectedDate string    `gorm:"column:selected_date" json:"selectedDate" db:"selectedDate"`
	DayOfWeek    string    `gorm:"column:dayofweek" json:"dayofweek" db:"dayofweek"`
	Time1        string    `gorm:"column:time1" json:"time1" db:"time1"`
	PartyRoom    Partyroom `gorm:"foreignKey:PartyRoomID;references:ID;onDelete:CASCADE"`
}

// 테이블 이름을 명시적으로 설정
func (SelectTime) TableName() string {
	return "selectTime" // 실제 테이블 이름
}
