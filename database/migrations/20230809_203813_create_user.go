package migrations

import (
	"github.com/atom-apps/door/common/consts"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_203813CreateUser) table() interface{} {
	type User struct {
		gorm.Model
		Uuid          string `gorm:"size:128"`
		Username      string `gorm:"size:128"`
		Password      string `gorm:"size:128"`
		Email         string `gorm:"size:128"`
		EmailVerified bool
		Phone         string `gorm:"size:128"`
		DisplayName   string `gorm:"size:128"`
		Avatar        string
		Status        consts.UserStatus `gorm:"size:64"`
	}

	return &User{}
}

func (m *Migration20230809_203813CreateUser) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_203813CreateUser) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_203813CreateUserMigration)
}

type Migration20230809_203813CreateUser struct {
	id string
}

func New20230809_203813CreateUserMigration() contracts.Migration {
	return &Migration20230809_203813CreateUser{id: "20230809_203813_create_user"}
}

func (m *Migration20230809_203813CreateUser) ID() string {
	return m.id
}
