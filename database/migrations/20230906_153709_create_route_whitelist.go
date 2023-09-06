package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230906_153709CreateRouteWhitelist) table() interface{} {
	type RouteWhitelist struct {
		ModelOnlyID
		Route string `gorm:"type:varchar(255);not null;comment:路由"`
	}

	return &RouteWhitelist{}
}

func (m *Migration20230906_153709CreateRouteWhitelist) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230906_153709CreateRouteWhitelist) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230906_153709CreateRouteWhitelistMigration)
}

type Migration20230906_153709CreateRouteWhitelist struct {
	id string
}

func New20230906_153709CreateRouteWhitelistMigration() contracts.Migration {
	return &Migration20230906_153709CreateRouteWhitelist{id: "20230906_153709_create_route_whitelist"}
}

func (m *Migration20230906_153709CreateRouteWhitelist) ID() string {
	return m.id
}
