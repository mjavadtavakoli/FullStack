package models

type MeetingUser struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	MeetingID uint
}
