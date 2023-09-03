package seeders

import (
	"github.com/atom-apps/door/database/models"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type PermissionsSeeder struct{}

func NewPermissionsSeeder() contracts.Seeder {
	return &PermissionsSeeder{}
}

func (s *PermissionsSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Permission{}).TableName(nil))
	items := []models.Permission{
		{TenantID: 1, RoleID: 2, Path: "/v1/users/tenants", Action: "GET"},
		{TenantID: 2, RoleID: 2, Path: "/v1/users/tenants", Action: "GET"},
	}
	lo.ForEach(items, func(m models.Permission, _ int) {
		db.Create(&m)
	})
}

func (s *PermissionsSeeder) Generate(faker *gofakeit.Faker, idx int) models.Permission {
	return models.Permission{}
}
