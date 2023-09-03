package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230903_141101CreatePermissions) table() interface{} {
	type Permissions struct {
		ModelOnlyID
		TenantID uint   `gorm:"not null;comment:租户ID"`
		RoleId   uint   `gorm:"not null;comment:角色ID"`
		Path     string `gorm:"varchar(256);not null;comment:路由"`
		Action   string `gorm:"varchar(24);not null;comment:请求方式"`
	}

	return &Permissions{}
}

func (m *Migration20230903_141101CreatePermissions) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230903_141101CreatePermissions) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230903_141101CreatePermissionsMigration)
}

type Migration20230903_141101CreatePermissions struct {
	id string
}

func New20230903_141101CreatePermissionsMigration() contracts.Migration {
	return &Migration20230903_141101CreatePermissions{id: "20230903_141101_create_permissions"}
}

func (m *Migration20230903_141101CreatePermissions) ID() string {
	return m.id
}
