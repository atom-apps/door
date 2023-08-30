package migrations

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_204703CreateUserInfo) table() interface{} {
	type UserInfo struct {
		Model

		UserID      uint          `gorm:"uniqueIndex; comment:用户ID"`
		Affiliation string        `gorm:"size:128; comment:工作单位"`
		Title       string        `gorm:"size:128; comment:职称"`
		IdCardType  string        `gorm:"size:128; not null; default ''; comment:证件类型"`
		IdCard      string        `gorm:"size:128; not null; default ''; comment:证件号码"`
		Biography   string        `gorm:"size:256; comment:自我介绍"`
		Tag         string        `gorm:"size:128; comment:标签"`
		Language    string        `gorm:"size:128; comment:语言"`
		Gender      consts.Gender `gorm:"not null; default ''; comment:性别"`
		Birthday    time.Time     `gorm:"comment:生日"`
		Education   string        `gorm:"size:128; comment:学历"`
		RealName    string        `gorm:"size:128; comment:真实姓名"`
	}

	return &UserInfo{}
}

func (m *Migration20230809_204703CreateUserInfo) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_204703CreateUserInfo) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_204703CreateUserInfoMigration)
}

type Migration20230809_204703CreateUserInfo struct {
	id string
}

func New20230809_204703CreateUserInfoMigration() contracts.Migration {
	return &Migration20230809_204703CreateUserInfo{id: "20230809_204703_create_user_info"}
}

func (m *Migration20230809_204703CreateUserInfo) ID() string {
	return m.id
}
