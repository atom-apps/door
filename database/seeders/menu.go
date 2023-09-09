package seeders

import (
	"github.com/atom-apps/door/database/models"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type MenuSeeder struct {
	id uint64
}

func (s *MenuSeeder) getID() uint64 {
	s.id++
	return s.id
}

func NewMenuSeeder() contracts.Seeder {
	return &MenuSeeder{}
}

func (s *MenuSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.PermissionRule{}).TableName(nil))

	models := []models.Menu{
		{ID: s.getID(), Name: "GroupMenu-1", Slug: "group_menu", GroupID: 0, ParentID: 0},
		{ID: s.getID(), Name: "GroupMenu-2", Slug: "group_menu_2", GroupID: 0, ParentID: 0},
		{ID: s.getID(), Name: "GroupMenu-1-Sub-1", Slug: "group_menu_2_1", GroupID: 1, ParentID: 1},
		{ID: s.getID(), Name: "GroupMenu-1-Sub-2", Slug: "group_menu_2_2", GroupID: 1, ParentID: 1},
		{ID: s.getID(), Name: "GroupMenu-1-Sub-3", Slug: "group_menu_2_3", GroupID: 1, ParentID: 3},
		{ID: s.getID(), Name: "GroupMenu-1-Sub-4", Slug: "group_menu_2_4", GroupID: 1, ParentID: 4},
	}

	db.CreateInBatches(models, 100)
}
