package seeders

import (
	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type TenantSeeder struct{}

func NewTenantSeeder() contracts.Seeder {
	return &TenantSeeder{}
}

func (s *TenantSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Tenant{}).TableName(nil))
	tenants := []models.Tenant{
		{
			Name:        "租户1",
			Description: "租户1",
			Meta:        "",
		},
		{
			Name:        "租户2",
			Description: "租户2",
			Meta:        "",
		},
	}

	db.CreateInBatches(tenants, len(tenants))
}

func (s *TenantSeeder) Generate(faker *gofakeit.Faker, idx int) models.Tenant {
	return models.Tenant{
		// fill model fields
	}
}
