package models

import "time"

type User struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"size:100" json:"name"`
	Email         string `gorm:"size:100;uniqueIndex" json:"email"`
	PasswordHash  string `gorm:"size:255" json:"password_hash"`
	EmailVerified bool   `gorm:"default:false" json:"email_verified"`

	Tasks []Task `gorm:"foreignKey:AssigneeID" json:"tasks"` // One-to-many relationship

	Organizations []UserOrganization // Many-to-many relationship through UserOrganization

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
