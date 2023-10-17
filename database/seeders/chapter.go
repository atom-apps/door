package seeders

import (
	"log"

	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type ChapterSeeder struct{}

func NewChapterSeeder() contracts.Seeder {
	return &ChapterSeeder{}
}

func (s *ChapterSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Chapter{}).TableName(nil))

	times := 200
	for i := 0; i < times; i++ {
		data := s.Generate(faker, i)
		if i == 0 {
			stmt := &gorm.Statement{DB: db}
			_ = stmt.Parse(&data)
			log.Printf("seeding %s for %d times", stmt.Schema.Table, times)
		}
		db.Create(&data)
	}
}

func (s *ChapterSeeder) Generate(faker *gofakeit.Faker, idx int) models.Chapter {
	return models.Chapter{
		TenantID:    1,
		UserID:      1,
		BookID:      uint64(faker.Number(1, 50)),
		Title:       faker.Book().Title,
		Description: faker.Book().Genre,
		Content:     faker.Paragraph(5, 10, 100, ","),
	}
}
