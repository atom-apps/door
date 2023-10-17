package seeders

import (
	"time"

	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/common/model"
	"github.com/atom-apps/door/database/models"
	"github.com/samber/lo"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/rogeecn/atom/contracts"
	dbUtil "github.com/rogeecn/atom/utils/db"
	"gorm.io/gorm"
)

type ArticleSeeder struct{}

func NewArticleSeeder() contracts.Seeder {
	return &ArticleSeeder{}
}

func (s *ArticleSeeder) Run(faker *gofakeit.Faker, db *gorm.DB) {
	dbUtil.TruncateTable(db, (&models.Article{}).TableName(nil))
	dbUtil.TruncateTable(db, (&models.ArticleAttachment{}).TableName(nil))
	dbUtil.TruncateTable(db, (&models.ArticleDig{}).TableName(nil))
	dbUtil.TruncateTable(db, (&models.ArticleContent{}).TableName(nil))
	dbUtil.TruncateTable(db, (&models.ArticlePayment{}).TableName(nil))
	dbUtil.TruncateTable(db, (&models.ArticleForwardSource{}).TableName(nil))

	times := 100
	ms := lo.Times(times, func(index int) *models.Article {
		data := s.Generate(faker, index, true)
		return &data
	})
	db.CreateInBatches(&ms, 100)

	ms = lo.Times(times, func(index int) *models.Article {
		data := s.Generate(faker, index, false)
		return &data
	})
	db.CreateInBatches(&ms, 100)

	times = 200

	// attachments
	attachments := lo.Times(times, func(idx int) *models.ArticleAttachment {
		data := s.GenerateAttachment(faker, idx+1)
		return &data
	})
	db.CreateInBatches(&attachments, 100)

	// digs
	digs := lo.Times(times, func(idx int) *models.ArticleDig {
		data := s.GenerateDig(faker, idx+1)
		return &data
	})
	db.CreateInBatches(&digs, 100)

	// contents
	contents := lo.Times(times, func(idx int) *models.ArticleContent {
		data := s.GenerateContent(faker, idx+1)
		return &data
	})
	db.CreateInBatches(&contents, 100)

	// payments
	payments := lo.Times(times, func(idx int) *models.ArticlePayment {
		data := s.GeneratePayment(faker, idx+1)
		return &data
	})
	db.CreateInBatches(&payments, 100)

	// forwards
	forwards := lo.Times(times, func(idx int) *models.ArticleForwardSource {
		data := s.GenerateForward(faker, idx+1)
		return &data
	})
	db.CreateInBatches(&forwards, 100)

	// tags
	lo.Times(5, func(index int) int {
		tags := lo.Times(times, func(idx int) *models.ArticleTag {
			data := s.GenerateTag(faker, idx+1)
			return &data
		})
		db.CreateInBatches(&tags, 100)
		return index
	})
}

func (s *ArticleSeeder) GenerateTag(faker *gofakeit.Faker, idx int) models.ArticleTag {
	return models.ArticleTag{
		ArticleID: uint64(idx),
		TagID:     uint64(faker.Number(1, 100)),
	}
}

func (s *ArticleSeeder) GenerateForward(faker *gofakeit.Faker, idx int) models.ArticleForwardSource {
	return models.ArticleForwardSource{
		CreatedAt:    time.Time{},
		ArticleID:    uint64(idx),
		Source:       faker.URL(),
		SourceAuthor: faker.Username(),
	}
}

func (s *ArticleSeeder) GeneratePayment(faker *gofakeit.Faker, idx int) models.ArticlePayment {
	return models.ArticlePayment{
		ArticleID:       uint64(idx),
		PriceType:       consts.ArticlePriceTypeContent,
		Token:           faker.Password(true, false, true, false, false, 4),
		Price:           uint64(faker.Price(0, 100)),
		Discount:        uint64(faker.Number(0, 100)),
		DiscountStartAt: faker.FutureDate(),
		DiscountEndAt:   faker.FutureDate(),
	}
}

func (s *ArticleSeeder) GenerateContent(faker *gofakeit.Faker, idx int) models.ArticleContent {
	return models.ArticleContent{
		TenantID:     1,
		UserID:       1,
		ArticleID:    uint64(idx),
		FreeContent:  faker.Paragraph(2, 4, 6, "\n"),
		PriceContent: faker.Paragraph(2, 4, 6, "\n"),
	}
}

func (s *ArticleSeeder) GenerateDig(faker *gofakeit.Faker, idx int) models.ArticleDig {
	return models.ArticleDig{
		ArticleID: uint64(idx),
		Views:     uint64(faker.Number(1, 100)),
		Likes:     uint64(faker.Number(1, 100)),
		Dislikes:  uint64(faker.Number(1, 100)),
	}
}

func (s *ArticleSeeder) GenerateAttachment(faker *gofakeit.Faker, idx int) models.ArticleAttachment {
	return models.ArticleAttachment{
		TenantID:     1,
		UserID:       1,
		ArticleID:    uint64(idx),
		FilesystemID: 1,
		Description:  faker.StreetName(),
		Password:     faker.Password(true, false, true, false, false, 4),
	}
}

func (s *ArticleSeeder) Generate(faker *gofakeit.Faker, idx int, isBook bool) models.Article {
	bookID := 0
	if isBook {
		bookID = faker.Number(1, 50)
	}
	chapterID := 0
	if isBook {
		chapterID = faker.Number(1, 200)
	}

	return models.Article{
		TenantID:    1,
		UserID:      1,
		UUID:        faker.UUID(),
		BookID:      uint64(bookID),
		ChapterID:   uint64(chapterID),
		CategoryID:  faker.Uint64() % 10,
		PublishAt:   faker.FutureDate(),
		Type:        consts.ArticleTypeValues()[faker.Number(0, len(consts.ArticleTypeValues())-1)],
		Format:      consts.ArticleFormatValues()[faker.Number(0, len(consts.ArticleFormatValues())-1)],
		Title:       faker.Book().Title,
		Keyword:     faker.StreetName(),
		Description: faker.StreetName(),
		Thumbnails: model.ArticleThumbnails{
			{Width: 100, Height: 100, Image: faker.ImageURL(100, 100), Head: false},
			{Width: 100, Height: 100, Image: faker.ImageURL(100, 100), Head: false},
			{Width: 100, Height: 100, Image: faker.ImageURL(100, 100), Head: false},
		},
		Videos: model.ArticleVideos{
			{Provider: "", Url: ""},
		},
		Audios: model.ArticleAudios{
			{Provider: "", Url: ""},
		},
		PostIP: faker.IPv4Address(),
		Weight: faker.Uint64(),
	}
}
