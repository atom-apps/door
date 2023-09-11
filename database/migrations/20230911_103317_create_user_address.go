package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230911_103317CreateUserAddress) table() interface{} {
	type UserAddress struct {
		Model
		UserID    uint   `gorm:"comment:用户ID"`
		Code      string `gorm:"size:6;comment:行政区划代码"`
		Town      string `gorm:"size:6;comment:街道"`
		Detail    string `gorm:"size:256;comment:详细地址"`
		Name      string `gorm:"size:64;comment:姓名"`
		Phone     string `gorm:"size:24;comment:联系电话"`
		ZipCode   string `gorm:"size:12;comment:邮编"`
		IsDefault bool   `gorm:"comment:默认地址"`
	}

	return &UserAddress{}
}

func (m *Migration20230911_103317CreateUserAddress) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230911_103317CreateUserAddress) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230911_103317CreateUserAddressMigration)
}

type Migration20230911_103317CreateUserAddress struct {
	id string
}

func New20230911_103317CreateUserAddressMigration() contracts.Migration {
	return &Migration20230911_103317CreateUserAddress{id: "20230911_103317_create_user_address"}
}

func (m *Migration20230911_103317CreateUserAddress) ID() string {
	return m.id
}
