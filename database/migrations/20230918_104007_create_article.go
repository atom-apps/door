package migrations

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/common/model"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_104007CreateArticle) table() interface{} {
	type Article struct {
		Model
		ModelWithUser
		UUID        string                  `gorm:"index:idx_uuid;size:64;not null;comment:UUID"`
		BookID      uint                    `gorm:"index;comment:书ID"`
		ChapterID   uint                    `gorm:"index;comment:章节ID"`
		CategoryID  uint                    `gorm:"index;comment:分类"`
		PublishAt   time.Time               `gorm:"comment:发布时间"`                                      // 未到达发布时间不发布
		Type        consts.ArticleType      `gorm:"type:int(11);index:idx_type;not null;comment:文章类型"` // 文字，图片，视频，音频
		Format      consts.ArticleFormat    `gorm:"type:int(11);not null;comment:文章格式"`                // html, markdown
		Title       string                  `gorm:"index:idx_title;size:128;not null;default:'';comment:标题"`
		Keyword     string                  `gorm:"size:256;not null;default:'';comment:关键词"`
		Description string                  `gorm:"size:256;not null;default:'';comment:简介"`
		Thumbnails  model.ArticleThumbnails `gorm:"type:text;not null;comment:缩略图"`
		Videos      model.ArticleVideos     `gorm:"type:text;not null;comment:视频"`
		Audios      model.ArticleAudios     `gorm:"type:text;not null;comment:音频"`
		PostIP      string                  `gorm:"size:128;comment:发布IP"`
		Weight      uint                    `gorm:"comment:权重"` // 0 - uint.max 最低 - 最高
	}

	return &Article{}
}

func (m *Migration20230918_104007CreateArticle) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_104007CreateArticle) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230918_104007CreateArticleMigration)
}

type Migration20230918_104007CreateArticle struct {
	id string
}

func New20230918_104007CreateArticleMigration() contracts.Migration {
	return &Migration20230918_104007CreateArticle{id: "20230918_104007_create_article"}
}

func (m *Migration20230918_104007CreateArticle) ID() string {
	return m.id
}
