package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_104216CreateArticleContent) table() interface{} {
	type ArticleContent struct {
		Model
		ModelWithUser
		ModelWithArticle
		FreeContent  string `gorm:"type:text;comment:内容"`
		PriceContent string `gorm:"type:text;comment:付费内容"`
	}

	return &ArticleContent{}
}

func (m *Migration20230918_104216CreateArticleContent) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_104216CreateArticleContent) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230918_104216CreateArticleContentMigration)
}

type Migration20230918_104216CreateArticleContent struct {
	id string
}

func New20230918_104216CreateArticleContentMigration() contracts.Migration {
	return &Migration20230918_104216CreateArticleContent{id: "20230918_104216_create_article_content"}
}

func (m *Migration20230918_104216CreateArticleContent) ID() string {
	return m.id
}
