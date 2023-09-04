package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230825_171341CreateRoutes) table() interface{} {
	type Routes struct {
		ModelOnlyID
		Type     string `gorm:"size:64;not null;comment:类型"`
		ParentID uint   `gorm:"not null;default 0;comment:父级ID"`
		Name     string `gorm:"size:255;not null;comment:名称"`
		Method   string `gorm:"size:24;not null;comment:请求方法"`
		Path     string `gorm:"size:1024;not null; comment:路径"`
		Metadata string `gorm:"default '{}'; comment:元数据"`
		Order    uint   `gorm:"default 0; comment:排序"`
	}

	return &Routes{}
}

func (m *Migration20230825_171341CreateRoutes) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230825_171341CreateRoutes) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230825_171341CreateRoutesMigration)
}

type Migration20230825_171341CreateRoutes struct {
	id string
}

func New20230825_171341CreateRoutesMigration() contracts.Migration {
	return &Migration20230825_171341CreateRoutes{id: "20230825_171341_create_routes"}
}

func (m *Migration20230825_171341CreateRoutes) ID() string {
	return m.id
}
