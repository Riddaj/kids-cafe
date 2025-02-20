package models

import (
	"time"
)

type SelectTime struct {
	ID          string    `json:"id" db:"id"`
	PartyRoomID string    `json:"partyroom_id" db:"partyroom_id"`
	Date        time.Time `json:"date" db:"date"`
	DayOfWeek   string    `json:"day_of_week" db:"dayofweek"`
	StartTime   time.Time `json:"start_time" db:"start_time"`
	EndTime     time.Time `json:"end_time" db:"end_time"`
	PartyRoom   Partyroom `gorm:"foreignKey:PartyRoomID;references:ID;onDelete:CASCADE"`
}
