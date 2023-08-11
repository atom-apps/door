package migrations

import (
	"time"

	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_203822CreateToken) table() interface{} {
	type Token struct {
		gorm.Model

		UserID        uint
		SessionID     uint
		AccessToken   string
		RefreshToken  string
		ExpireAt      time.Time
		Scope         string `gorm:"size:128"`
		TokenType     string `gorm:"size:128"`
		CodeChallenge string `gorm:"size:128"`
		Code          string `gorm:"unique"`
		CodeExpireAt  time.Time
		Used          bool
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
