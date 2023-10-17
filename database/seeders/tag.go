package seeders

import (
	"github.com/atom-apps/door/database/models"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type TagSeeder struct{}

func NewTagSeeder() contracts.Seeder {
	return &TagSeeder{}
}

func (s *TagSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Tag{}).TableName(nil))

	times := 100
	ms := lo.Times(times, func(index int) *models.Tag {
		item := s.Generate(faker, index)
		return &item
	})

	db.CreateInBatches(ms, 100)
}

func (s *TagSeeder) Generate(faker *gofakeit.Faker, idx int) models.Tag {
	name := faker.Name()
	if len(name) > 12 {
		name = name[0:12]
	}
	return models.Tag{
		Name:  name,
		Count: uint64(faker.Number(1, 10000)),
	}
}
