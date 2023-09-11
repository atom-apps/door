package seeders

import (
	"encoding/json"
	"log"

	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"github.com/rogeecn/fabfile"
	"gorm.io/gorm"
)

type LocationSeeder struct{}

func NewLocationSeeder() contracts.Seeder {
	return &LocationSeeder{}
}

func (s *LocationSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Location{}).TableName(nil))

	location := fabfile.MustRead("location.json")

	var locations []*models.Location
	if err := json.Unmarshal(location, &locations); err != nil {
		log.Fatal(err)
	}

	db.CreateInBatches(locations, 100)
}
