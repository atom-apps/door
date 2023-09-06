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

type UserSeeder struct{}

func NewUserSeeder() contracts.Seeder {
	return &UserSeeder{}
}

func (s *UserSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.User{}).TableName(nil))

	users := []models.User{
		{
			UUID:          "3b680cc7-b823-4f27-8f90-3e7e1f5ddd55",
			Username:      "admin",
			Password:      "$2a$10$kXfWmId6tVAKkC5XI.fTluw2t3QEhwBN0otxKTXdRB/1G7lMOSnXG", // admin
			Email:         "",
			EmailVerified: true,
			Phone:         "18601010101",
			DisplayName:   "Admin",
			Avatar:        "",
			Status:        consts.UserStatusDefault,
		},
		{
			UUID:          "554687ef-39e1-4cb8-a4be-a16ff920b4ac",
			Username:      "system_admin",
			Password:      "$2a$10$kXfWmId6tVAKkC5XI.fTluw2t3QEhwBN0otxKTXdRB/1G7lMOSnXG", // admin
			Email:         "",
			EmailVerified: true,
			Phone:         "18601010101",
			DisplayName:   "SystemAdmin",
			Avatar:        "",
			Status:        consts.UserStatusDefault,
		},
		{
			UUID:          "554687ef-39e1-4cb8-a4be-a16ff920b1ac",
			Username:      "tenant_admin",
			Password:      "$2a$10$kXfWmId6tVAKkC5XI.fTluw2t3QEhwBN0otxKTXdRB/1G7lMOSnXG", // admin
			Email:         "",
			EmailVerified: true,
			Phone:         "18601010101",
			DisplayName:   "TenantAdmin",
			Avatar:        "",
			Status:        consts.UserStatusDefault,
		},
		{
			UUID:          "4e3d0fd6-a36c-4813-b5a6-005b3b1d911d",
			Username:      "tenant_user",
			Password:      "$2a$10$kXfWmId6tVAKkC5XI.fTluw2t3QEhwBN0otxKTXdRB/1G7lMOSnXG", // admin
			Email:         "",
			EmailVerified: true,
			Phone:         "18601010101",
			DisplayName:   "TenantUser",
			Avatar:        "",
			Status:        consts.UserStatusDefault,
		},
	}

	lo.ForEach(users, func(user models.User, _ int) {
		db.Create(&user)
	})
}

func (s *UserSeeder) Generate(faker *gofakeit.Faker, idx int) models.User {
	return models.User{
		// fill model fields
	}
}
