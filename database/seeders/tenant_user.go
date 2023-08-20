package seeders

import (
	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type TenantUserSeeder struct{}

func NewTenantUserSeeder() contracts.Seeder {
	return &TenantUserSeeder{}
}

func (s *TenantUserSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.TenantUser{}).TableName(nil))
	db.CreateInBatches(s.Generate(faker, 1), 10)
}

func (s *TenantUserSeeder) Generate(faker *gofakeit.Faker, idx int) []models.TenantUser {
	return []models.TenantUser{
		{TenantID: 1, UserID: 2},
		{TenantID: 2, UserID: 3},
	}
}
