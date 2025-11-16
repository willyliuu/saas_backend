package models

import (
	"time"
)

type Invite struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Email string `gorm:"size:100" json:"email"`
	Role  string `gorm:"size:50" json:"role"`

	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"organization"`

	Token     string    `gorm:"size:255" json:"token"`
	ExpiresAt time.Time `json:"expires_at"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
