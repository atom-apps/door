package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_210420CreateTenant) table() interface{} {
	type Tenant struct {
		ID          uint   `gorm:"primarykey"`
		Name        string `gorm:"size:64"`
		Description string `gorm:"size:256"`
		Meta        string
	}

	return &Tenant{}
}

func (m *Migration20230809_210420CreateTenant) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_210420CreateTenant) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_210420CreateTenantMigration)
}

type Migration20230809_210420CreateTenant struct {
	id string
}

func New20230809_210420CreateTenantMigration() contracts.Migration {
	return &Migration20230809_210420CreateTenant{id: "20230809_210420_create_tenant"}
}

func (m *Migration20230809_210420CreateTenant) ID() string {
	return m.id
}
