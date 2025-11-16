package models

import (
	"time"
)

type Organization struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100" json:"name"`

	OwnerID uint `json:"owner_id"`

	Users    []UserOrganization // Many-to-many relationship through UserOrganization
	Invites  []Invite           // One-to-many relationship
	Projects []Project          // One-to-many relationship

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
