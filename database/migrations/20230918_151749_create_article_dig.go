package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_151749CreateArticleDig) table() interface{} {
	type ArticleDig struct {
		ModelOnlyID
		ModelWithArticle
		Views    uint `gorm:"comment:浏览量"`
		Likes    uint `gorm:"comment:喜欢"`
		Dislikes uint `gorm:"comment:不喜欢"`
	}

	return &ArticleDig{}
}

func (m *Migration20230918_151749CreateArticleDig) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_151749CreateArticleDig) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230918_151749CreateArticleDigMigration)
}

type Migration20230918_151749CreateArticleDig struct {
	id string
}

func New20230918_151749CreateArticleDigMigration() contracts.Migration {
	return &Migration20230918_151749CreateArticleDig{id: "20230918_151749_create_article_dig"}
}

func (m *Migration20230918_151749CreateArticleDig) ID() string {
	return m.id
}
