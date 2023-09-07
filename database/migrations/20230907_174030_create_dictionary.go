package migrations

import (
	"github.com/atom-apps/door/common"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230907_174030CreateDictionary) table() interface{} {
	type Dictionary struct {
		Model
		Name        string            `gorm:"type:varchar(64);not null;comment:名称"`
		Slug        string            `gorm:"type:varchar(64);not null;comment:别名"`
		Description string            `gorm:"type:varchar(198);not null;comment:描述"`
		Items       common.LabelItems `gorm:"not null;comment:选项"`
	}

	return &Dictionary{}
}

func (m *Migration20230907_174030CreateDictionary) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230907_174030CreateDictionary) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230907_174030CreateDictionaryMigration)
}

type Migration20230907_174030CreateDictionary struct {
	id string
}

func New20230907_174030CreateDictionaryMigration() contracts.Migration {
	return &Migration20230907_174030CreateDictionary{id: "20230907_174030_create_dictionary"}
}

func (m *Migration20230907_174030CreateDictionary) ID() string {
	return m.id
}
