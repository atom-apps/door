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
		{Name: "管理组", Description: "后台管理用户组", Meta: ""},
		{Name: "用户组", Description: "普通注册用户", Meta: ""},
	}

	db.CreateInBatches(tenants, len(tenants))
}

func (s *TenantSeeder) Generate(faker *gofakeit.Faker, idx int) models.Tenant {
	return models.Tenant{
		// fill model fields
	}
}
