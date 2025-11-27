package migrations

import (
	"gin/internal/database"
	"gin/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrations() *gormigrate.Gormigrate {
	m := gormigrate.New(database.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20251116001_create_users_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.User{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("users")
			},
		},
		{
			ID: "20251116002_create_organizations_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Organization{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("organizations")
			},
		},
		{
			ID: "20251116003_create_user_organizations_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.UserOrganization{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("user_organizations")
			},
		},
		{
			ID: "20251116004_create_invites_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Invite{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("invites")
			},
		},
		{
			ID: "20251116005_create_projects_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Project{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("projects")
			},
		},
		{
			ID: "20251116006_create_tasks_table",
			Migrate: func(d *gorm.DB) error {
				return d.AutoMigrate(&models.Task{})
			},
			Rollback: func(d *gorm.DB) error {
				return d.Migrator().DropTable("tasks")
			},
		},
	})

	return m
}
