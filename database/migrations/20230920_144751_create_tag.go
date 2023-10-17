package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230920_144751CreateTag) table() interface{} {
	type Tag struct {
		Model
		Name  string `gorm:"size:12;not null;comment:名称"`
		Count uint   `gorm:"comment:引用次数"`
		Views uint   `gorm:"comment:点击次数"`
	}

	return &Tag{}
}

func (m *Migration20230920_144751CreateTag) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230920_144751CreateTag) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230920_144751CreateTagMigration)
}

type Migration20230920_144751CreateTag struct {
	id string
}

func New20230920_144751CreateTagMigration() contracts.Migration {
	return &Migration20230920_144751CreateTag{id: "20230920_144751_create_tag"}
}

func (m *Migration20230920_144751CreateTag) ID() string {
	return m.id
}
