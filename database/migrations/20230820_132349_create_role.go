package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230820_132349CreateRole) table() interface{} {
	type Role struct {
		ID          uint   `gorm:"primarykey"`
		Name        string `gorm:"size:128"`
		Slug        string `gorm:"size:128"`
		Description string `gorm:"size:256"`
		ParentID    uint
	}

	return &Role{}
}

func (m *Migration20230820_132349CreateRole) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230820_132349CreateRole) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230820_132349CreateRoleMigration)
}

type Migration20230820_132349CreateRole struct {
	id string
}

func New20230820_132349CreateRoleMigration() contracts.Migration {
	return &Migration20230820_132349CreateRole{id: "20230820_132349_create_role"}
}

func (m *Migration20230820_132349CreateRole) ID() string {
	return m.id
}
