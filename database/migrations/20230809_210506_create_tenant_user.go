package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_210506CreateTenantUser) table() interface{} {
	type TenantUser struct {
		gorm.Model
		TenantID int
		UserID   int
		IsAdmin  bool
	}

	return &TenantUser{}
}

func (m *Migration20230809_210506CreateTenantUser) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_210506CreateTenantUser) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_210506CreateTenantUserMigration)
}

type Migration20230809_210506CreateTenantUser struct {
	id string
}

func New20230809_210506CreateTenantUserMigration() contracts.Migration {
	return &Migration20230809_210506CreateTenantUser{id: "20230809_210506_create_tenant_user"}
}

func (m *Migration20230809_210506CreateTenantUser) ID() string {
	return m.id
}
