package models

import (
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:100" json:"title"`
	Description string    `gorm:"size:255" json:"description"`
	Status      string    `gorm:"size:50" json:"status"`
	DueDate     time.Time `json:"due_date"`

	ProjectID uint    `json:"project_id"`
	Project   Project `gorm:"foreignKey:ProjectID;references:ID;constraint:OnDelete:CASCADE" json:"project"`

	AssigneeID uint `json:"assignee_id"`
	Assignee   User `gorm:"foreignKey:AssigneeID;references:ID;constraint:OnDelete:SET NULL" json:"assignee"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
