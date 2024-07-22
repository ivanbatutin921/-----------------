package models

type User struct {
	ID       int    `json:"id" gorm:"type:uuid;primary_key;autoIncrement"`
	Login    string `json:"name" gorm:"not null"`
	Password string `json:"email" gorm:"not null"`
}
