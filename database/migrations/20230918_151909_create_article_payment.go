package migrations

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/rogeecn/atom/contracts"
	"gorm.io/gorm"
)

func (m *Migration20230918_151909CreateArticlePayment) table() interface{} {
	type ArticlePayment struct {
		ModelOnlyID
		ModelWithArticle

		PriceType       consts.ArticlePriceType `gorm:"comment:付费类型"`
		Token           string                  `gorm:"size:12;comment:付费密码"`
		Price           uint                    `gorm:"comment:付费价格"`
		Discount        uint                    `gorm:"comment:付费折扣"`
		DiscountStartAt time.Time               `gorm:"comment:折扣开始时间"`
		DiscountEndAt   time.Time               `gorm:"comment:折扣结束时间"`
	}

	return &ArticlePayment{}
}

func (m *Migration20230918_151909CreateArticlePayment) Up(tx *gorm.DB) error {
	return tx.AutoMigrate(m.table())
}

func (m *Migration20230918_151909CreateArticlePayment) Down(tx *gorm.DB) error {
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
	Migrations = append(Migrations, New20230918_151909CreateArticlePaymentMigration)
}

type Migration20230918_151909CreateArticlePayment struct {
	id string
}

func New20230918_151909CreateArticlePaymentMigration() contracts.Migration {
	return &Migration20230918_151909CreateArticlePayment{id: "20230918_151909_create_article_payment"}
}

func (m *Migration20230918_151909CreateArticlePayment) ID() string {
	return m.id
}
