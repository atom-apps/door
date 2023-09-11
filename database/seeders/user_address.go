package seeders

import (
	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type UserAddressSeeder struct{}

func NewUserAddressSeeder() contracts.Seeder {
	return &UserAddressSeeder{}
}

func (s *UserAddressSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.UserAddress{}).TableName(nil))
	ms := []*models.UserAddress{
		{ID: 1, UserID: 1, Code: "110101", Town: "001000", Detail: "景山前街4号", Name: "张三", Phone: "13800138000", ZipCode: "100000"},
	}
	db.CreateInBatches(ms, 100)
}

func (s *UserAddressSeeder) Generate(faker *gofakeit.Faker, idx int) models.UserAddress {
	return models.UserAddress{
		// fill model fields
	}
}
