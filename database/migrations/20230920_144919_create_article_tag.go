package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230920_144919CreateArticleTag) table() interface{} {
	type ArticleTag struct {
		ModelOnlyID
		ModelWithArticle
		TagID uint `gorm:"comment:标签ID"`
	}

	return &ArticleTag{}
}

func (m *Migration20230920_144919CreateArticleTag) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230920_144919CreateArticleTag) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230920_144919CreateArticleTagMigration)
}

type Migration20230920_144919CreateArticleTag struct {
	id string
}

func New20230920_144919CreateArticleTagMigration() contracts.Migration {
	return &Migration20230920_144919CreateArticleTag{id: "20230920_144919_create_article_tag"}
}

func (m *Migration20230920_144919CreateArticleTag) ID() string {
	return m.id
}
