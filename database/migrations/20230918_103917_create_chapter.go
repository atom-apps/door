package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_103917CreateChapter) table() interface{} {
	type Chapter struct {
		Model
		ModelWithUser
		BookID      uint   `gorm:"comment:书ID"`
		Title       string `gorm:"size:128;not null;default:'';comment:章节名"`
		Description string `gorm:"size:256;not null;default:'';comment:简介"`
		Content     string `gorm:"type:text;comment:详细介绍"`
	}

	return &Chapter{}
}

func (m *Migration20230918_103917CreateChapter) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_103917CreateChapter) Down(tx *gorm.DB) error {
	return tx.Migrator().DropTable(m.table())
	// return tx.Migrator().DropColumn(m.table(), "input_column_name")
}

// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
// DO NOT EDIT BLOW CODES!!
func init() {
	Migrations = append(Migrations, New20230918_103917CreateChapterMigration)
}

type Migration20230918_103917CreateChapter struct {
	id string
}

func New20230918_103917CreateChapterMigration() contracts.Migration {
	return &Migration20230918_103917CreateChapter{id: "20230918_103917_create_chapter"}
}

func (m *Migration20230918_103917CreateChapter) ID() string {
	return m.id
}
