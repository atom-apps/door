package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_205909CreatePermissionRule) table() interface{} {
	type PermissionRule struct {
		ID    uint   `gorm:"primarykey"`
		Ptype string `gorm:"size:128; index; not null; default ''`
		V0    string `gorm:"size:128; index; not null; default ''`
		V1    string `gorm:"size:128; index; not null; default ''`
		V2    string `gorm:"size:128; index; not null; default ''`
		V3    string `gorm:"size:128; index; not null; default ''`
		V4    string `gorm:"size:128; index; not null; default ''`
		V5    string `gorm:"size:128; index; not null; default ''`
	}

	return &PermissionRule{}
}

func (m *Migration20230809_205909CreatePermissionRule) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_205909CreatePermissionRule) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_205909CreatePermissionRuleMigration)
}

type Migration20230809_205909CreatePermissionRule struct {
	id string
}

func New20230809_205909CreatePermissionRuleMigration() contracts.Migration {
	return &Migration20230809_205909CreatePermissionRule{id: "20230809_205909_create_permission_rule"}
}

func (m *Migration20230809_205909CreatePermissionRule) ID() string {
	return m.id
}
