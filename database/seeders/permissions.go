package seeders

import (
	"github.com/atom-apps/door/database/models"

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
		{TenantID: 1, RoleID: 3, RouteID: 2},
		{TenantID: 1, RoleID: 3, RouteID: 4},
		{TenantID: 1, RoleID: 3, RouteID: 29},
	}
	db.CreateInBatches(&items, 10)
}

func (s *PermissionsSeeder) Generate(faker *gofakeit.Faker, idx int) models.Permission {
	return models.Permission{}
}
