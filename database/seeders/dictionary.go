package seeders

import (
	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	dbUtil "github.com/rogeecn/atom/utils/db"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

type DictionarySeeder struct{}

func NewDictionarySeeder() contracts.Seeder {
	return &DictionarySeeder{}
}

func (s *DictionarySeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Dictionary{}).TableName(nil))
	items := []models.Dictionary{
		{
			Name:        "性别",
			Slug:        "gender",
			Description: "性别",
			Items: []ds.LabelItem{
				{Label: "未知", Value: "unknown"},
				{Label: "男", Value: "male"},
				{Label: "女", Value: "female"},
			},
		},
		{
			Name:        "状态",
			Slug:        "status",
			Description: "状态值表示",
			Items: []ds.LabelItem{
				{Label: "启用", Value: "enabled"},
				{Label: "禁用", Value: "disabled"},
			},
		},
	}

	db.CreateInBatches(items, 10)
}

func (s *DictionarySeeder) Generate(faker *gofakeit.Faker, idx int) models.Dictionary {
	return models.Dictionary{
		// fill model fields
	}
}
