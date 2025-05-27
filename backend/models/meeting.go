package models

import "time"

type Meeting struct {
	ID       uint      `gorm:"primaryKey"`
	Title    string    `json:"title"`
	DateTime time.Time `json:"date_time"`
	OwnerID  uint      `json:"owner_id"`
	Users    []User    `gorm:"many2many:meeting_users;"`
}
