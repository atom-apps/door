package seeders

import (
	"github.com/atom-apps/door/database/models"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type UserTenantRoleSeeder struct{}

func NewUserTenantRoleSeeder() contracts.Seeder {
	return &UserTenantRoleSeeder{}
}

func (s *UserTenantRoleSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.UserTenantRole{}).TableName(nil))
	items := []models.UserTenantRole{
		{UserID: 1, TenantID: 0, RoleID: 1},
		{UserID: 2, TenantID: 1, RoleID: 2},
		{UserID: 3, TenantID: 2, RoleID: 2},
		{UserID: 3, TenantID: 2, RoleID: 3},
	}
	lo.ForEach(items, func(m models.UserTenantRole, _ int) {
		db.Create(&m)
	})
}

func (s *UserTenantRoleSeeder) Generate(faker *gofakeit.Faker, idx int) models.UserTenantRole {
	return models.UserTenantRole{}
}
