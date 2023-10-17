package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_161744CreateArticleForwardSource) table() interface{} {
	type ArticleForwardSource struct {
		ModelOnlyID
		ModelWithArticle

		Source       string `gorm:"size:1024;comment:原文地址"`
		SourceAuthor string `gorm:"size:64;comment:原文作者"`
	}

	return &ArticleForwardSource{}
}

func (m *Migration20230918_161744CreateArticleForwardSource) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_161744CreateArticleForwardSource) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230918_161744CreateArticleForwardSourceMigration)
}

type Migration20230918_161744CreateArticleForwardSource struct {
	id string
}

func New20230918_161744CreateArticleForwardSourceMigration() contracts.Migration {
	return &Migration20230918_161744CreateArticleForwardSource{id: "20230918_161744_create_article_forward_source"}
}

func (m *Migration20230918_161744CreateArticleForwardSource) ID() string {
	return m.id
}
