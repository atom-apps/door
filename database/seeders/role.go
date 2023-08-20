package seeders

import (
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type RoleSeeder struct{}

func NewRoleSeeder() contracts.Seeder {
	return &RoleSeeder{}
}

func (s *RoleSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Role{}).TableName(nil))

	roles := []models.Role{
		{Name: "超级管理员", Slug: consts.RoleSuperAdmin.String(), Description: "超级管理员", ParentID: 0},
		{Name: "租户管理员", Slug: consts.RoleTenantAdmin.String(), Description: "租户管理员", ParentID: 1},
	}
	lo.ForEach(roles, func(role models.Role, _ int) {
		db.Create(&role)
	})
}

func (s *RoleSeeder) Generate(faker *gofakeit.Faker, idx int) models.Role {
	return models.Role{
		// fill model fields
	}
}
