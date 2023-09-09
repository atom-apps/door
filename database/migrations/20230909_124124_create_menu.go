package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230909_124124CreateMenu) table() interface{} {
	type Menu struct {
		ModelOnlyID
		Name     string `gorm:"size:64;comment:名称"`
		Slug     string `gorm:"size:128;comment:别名"`
		GroupID  uint   `gorm:"comment:组"`   // 0: 为菜单组的分类, 菜单是他的子级
		ParentID uint   `gorm:"comment:父ID"` // 顶级的不存在父ID
		Metadata string `gorm:"type:text;comment:元数据"`
	}

	return &Menu{}
}

func (m *Migration20230909_124124CreateMenu) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230909_124124CreateMenu) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230909_124124CreateMenuMigration)
}

type Migration20230909_124124CreateMenu struct {
	id string
}

func New20230909_124124CreateMenuMigration() contracts.Migration {
	return &Migration20230909_124124CreateMenu{id: "20230909_124124_create_menu"}
}

func (m *Migration20230909_124124CreateMenu) ID() string {
	return m.id
}
