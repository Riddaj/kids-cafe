package models

import (
	"time"

	"gorm.io/gorm"
)

// 1. 시간대 테이블 (예약 가능한 시간 슬롯 저장)
type TimeSlot struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StartTime time.Time `json:"start_time"` // ex) 10:00 AM
	EndTime   time.Time `json:"end_time"`   // ex) 12:00 PM
	Capacity  int       `json:"capacity"`   // 해당 시간대 최대 예약 가능 인원
}

// Service table (예: 입장권, party packages etc)
type Service struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name"`
	RoomType    string     `json:"roomtype"`
	DayType     string     `json:"daytype"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Location    string     `json:"location"`
	TimeSlots   []TimeSlot `gorm:"many2many:booking_time_slots;" json:"time_slots"` // 다대다 관계
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Booking struct {
	gorm.Model
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	KidsCount   int       `json:"kids_count"`
	AdultsCount int       `json:"adults_count"`
	Date        time.Time `json:"date"`
	Products    []Service `gorm:"many2many:booking_services;" json:"services"` // 다대다 관계

}

// Book-product 연결 테이블 (다대다 관계)
type BookingService struct {
	BookingID uint `gorm:"primaryKey"`
	ServiceID uint `gorm:"primaryKey"`
	Quantity  int  `json:"quantity"` // 선택한 상품 개수
}

// 3. 예약-시간대 연결 테이블 (다대다 관계)
type BookingTimeSlot struct {
	BookingID  uint `gorm:"primaryKey"`
	TimeSlotID uint `gorm:"primaryKey"`
}
