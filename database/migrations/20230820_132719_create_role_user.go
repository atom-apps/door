package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230820_132719CreateRoleUser) table() interface{} {
	type RoleUser struct {
		ID       uint `gorm:"primarykey"`
		RoleID   uint
		UserID   uint
		TenantID uint
	}

	return &RoleUser{}
}

func (m *Migration20230820_132719CreateRoleUser) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230820_132719CreateRoleUser) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230820_132719CreateRoleUserMigration)
}

type Migration20230820_132719CreateRoleUser struct {
	id string
}

func New20230820_132719CreateRoleUserMigration() contracts.Migration {
	return &Migration20230820_132719CreateRoleUser{id: "20230820_132719_create_role_user"}
}

func (m *Migration20230820_132719CreateRoleUser) ID() string {
	return m.id
}
