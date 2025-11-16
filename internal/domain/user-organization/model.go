package userorganization

import (
	"time"

	"gin/internal/domain/organization"
	"gin/internal/domain/user"
)

type UserOrganization struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint      `json:"user_id"`
	User   user.User `gorm:"foreignKey:UserID;references:ID" json:"user"`

	OrganizationID uint                      `json:"organization_id"`
	Organization   organization.Organization `gorm:"foreignKey:OrganizationID;references:ID" json:"organization"`

	Role      string    `gorm:"size:50" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
