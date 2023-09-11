package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230822_140959CreateLocation) table() interface{} {
	type Location struct {
		ModelOnlyID

		Code     string `gorm:"size:6;comment:行政区划代码"`
		Name     string `gorm:"size:128;comment:名称"`
		Province string `gorm:"size:2;comment:省/直辖市"`
		City     string `gorm:"size:2;comment:市"`
		Area     string `gorm:"size:2;comment:区县"`
		Town     string `gorm:"size:12;comment:乡镇"`
	}

	return &Location{}
}

func (m *Migration20230822_140959CreateLocation) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230822_140959CreateLocation) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230822_140959CreateLocationMigration)
}

type Migration20230822_140959CreateLocation struct {
	id string
}

func New20230822_140959CreateLocationMigration() contracts.Migration {
	return &Migration20230822_140959CreateLocation{id: "20230822_140959_create_location"}
}

func (m *Migration20230822_140959CreateLocation) ID() string {
	return m.id
}
