package migrations

import (
	"github.com/atom-apps/door/common/consts"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_163721CreateArticlePaidUser) table() interface{} {
	type ArticlePaidUser struct {
		ModelOnlyID
		ModelWithUser
		ModelWithArticle

		PriceType consts.ArticlePriceType `gorm:"comment:付费类型"`
		Price     uint                    `gorm:"comment:付费价格"`
		EndAt     uint                    `gorm:"comment:付费结束时间"`
	}

	return &ArticlePaidUser{}
}

func (m *Migration20230918_163721CreateArticlePaidUser) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_163721CreateArticlePaidUser) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230918_163721CreateArticlePaidUserMigration)
}

type Migration20230918_163721CreateArticlePaidUser struct {
	id string
}

func New20230918_163721CreateArticlePaidUserMigration() contracts.Migration {
	return &Migration20230918_163721CreateArticlePaidUser{id: "20230918_163721_create_article_paid_user"}
}

func (m *Migration20230918_163721CreateArticlePaidUser) ID() string {
	return m.id
}
