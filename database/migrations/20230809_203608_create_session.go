package migrations

import (
	"time"

	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230809_203608CreateSession) table() interface{} {
	type Session struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		UserID    uint
		SessionID string `gorm:"size:64"`
		ExpireAt  time.Time
	}

	return &Session{}
}

func (m *Migration20230809_203608CreateSession) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230809_203608CreateSession) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230809_203608CreateSessionMigration)
}

type Migration20230809_203608CreateSession struct {
	id string
}

func New20230809_203608CreateSessionMigration() contracts.Migration {
	return &Migration20230809_203608CreateSession{id: "20230809_203608_create_session"}
}

func (m *Migration20230809_203608CreateSession) ID() string {
	return m.id
}
