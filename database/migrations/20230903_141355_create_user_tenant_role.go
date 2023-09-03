package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230903_141355CreateUserTenantRole) table() interface{} {
	type UserTenantRole struct {
		ModelOnlyID
		UserID   uint `gorm:"not null;comment:用户ID"`
		TenantID uint `gorm:"not null;comment:租户ID"`
		RoleID   uint `gorm:"not null;comment:角色ID"`
	}

	return &UserTenantRole{}
}

func (m *Migration20230903_141355CreateUserTenantRole) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230903_141355CreateUserTenantRole) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230903_141355CreateUserTenantRoleMigration)
}

type Migration20230903_141355CreateUserTenantRole struct {
	id string
}

func New20230903_141355CreateUserTenantRoleMigration() contracts.Migration {
	return &Migration20230903_141355CreateUserTenantRole{id: "20230903_141355_create_user_tenant_role"}
}

func (m *Migration20230903_141355CreateUserTenantRole) ID() string {
	return m.id
}
