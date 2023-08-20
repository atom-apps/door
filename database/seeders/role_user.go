package seeders

import (
	"github.com/atom-apps/door/database/models"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type RoleUserSeeder struct{}

func NewRoleUserSeeder() contracts.Seeder {
	return &RoleUserSeeder{}
}

func (s *RoleUserSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.RoleUser{}).TableName(nil))

	ms := []models.RoleUser{
		{RoleID: 1, UserID: 1, TenantID: 0},
		{RoleID: 2, UserID: 2, TenantID: 1},
		{RoleID: 2, UserID: 3, TenantID: 2},
	}

	lo.ForEach(ms, func(model models.RoleUser, _ int) {
		db.Create(&model)
	})
}

func (s *RoleUserSeeder) Generate(faker *gofakeit.Faker, idx int) models.RoleUser {
	return models.RoleUser{
		// fill model fields
	}
}
