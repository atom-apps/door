package migrations

import (
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_105240CreateArticleAttachment) table() interface{} {
	type ArticleAttachment struct {
		Model
		ModelWithUser
		ModelWithArticle
		FilesystemID uint   `gorm:"comment:文件系统ID"`
		Description  string `gorm:"size:256;not null;default:'';comment:文件描述"`
		Password     string `gorm:"size:128;not null;default:'';comment:文件密码"`
	}

	return &ArticleAttachment{}
}

func (m *Migration20230918_105240CreateArticleAttachment) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_105240CreateArticleAttachment) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230918_105240CreateArticleAttachmentMigration)
}

type Migration20230918_105240CreateArticleAttachment struct {
	id string
}

func New20230918_105240CreateArticleAttachmentMigration() contracts.Migration {
	return &Migration20230918_105240CreateArticleAttachment{id: "20230918_105240_create_article_attachment"}
}

func (m *Migration20230918_105240CreateArticleAttachment) ID() string {
	return m.id
}
