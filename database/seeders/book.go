package seeders

import (
	"log"

	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type BookSeeder struct{}

func NewBookSeeder() contracts.Seeder {
	return &BookSeeder{}
}

func (s *BookSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Book{}).TableName(nil))

	times := 50
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

func (s *BookSeeder) Generate(faker *gofakeit.Faker, idx int) models.Book {
	return models.Book{
		TenantID:        1,
		UserID:          1,
		Title:           faker.Book().Title,
		ThumbnailFileID: uint64(faker.Number(1, 100)),
		Description:     faker.Street(),
		Content:         faker.Paragraph(10, 10, 10, ","),
		Author:          faker.Username(),
		Source:          faker.URL(),
		Isbn:            faker.MonthString(),
		Price:           uint64(faker.Price(10, 1000)),
	}
}
