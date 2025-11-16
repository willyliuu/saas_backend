package models

import (
	"time"
)

type Project struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100" json:"name"`
	Description string `gorm:"size:255" json:"description"`
	Archieved   bool   `gorm:"default:false" json:"archived"`

	OrganizationID uint         `json:"organization_id"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID;constraint:OnDelete:CASCADE" json:"organization"`

	Tasks []Task // One-to-many relationship

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
