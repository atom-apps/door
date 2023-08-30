package migrations

import (
	"time"

	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_203822CreateToken) table() interface{} {
	type Token struct {
		gorm.Model

		TenantID      uint      `gorm:"comment:租户ID"`
		UserID        uint      `gorm:"comment:用户ID"`
		SessionID     uint      `gorm:"comment:会话ID"`
		AccessToken   string    `gorm:"size:1024; comment:访问令牌"`
		RefreshToken  string    `gorm:"size:1024; comment:刷新令牌"`
		ExpireAt      time.Time `gorm:"comment:过期时间"`
		Scope         string    `gorm:"size:128; comment:Scope"`
		TokenType     string    `gorm:"size:128; comment:令牌类型"`
		CodeChallenge string    `gorm:"size:128; comment:CodeChallenge"`
		Code          string    `gorm:"unique; size:128; comment:Code"`
		CodeExpireAt  time.Time `gorm:"comment:Code过期时间"`
		Used          bool      `gorm:"comment:是否已使用"`
	}

	return &Token{}
}

func (m *Migration20230809_203822CreateToken) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_203822CreateToken) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_203822CreateTokenMigration)
}

type Migration20230809_203822CreateToken struct {
	id string
}

func New20230809_203822CreateTokenMigration() contracts.Migration {
	return &Migration20230809_203822CreateToken{id: "20230809_203822_create_token"}
}

func (m *Migration20230809_203822CreateToken) ID() string {
	return m.id
}
