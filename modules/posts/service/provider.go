package service

import (
	"github.com/atom-apps/door/modules/posts/dao"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		articleDao *dao.ArticleDao,
		attachmentDao *dao.ArticleAttachmentDao,
		contentDao *dao.ArticleContentDao,
		digDao *dao.ArticleDigDao,
		forwardSourceDao *dao.ArticleForwardSourceDao,
		paidUserDao *dao.ArticlePaidUserDao,
		paymentDao *dao.ArticlePaymentDao,
	) (*ArticleService, error) {
		obj := &ArticleService{
			articleDao:       articleDao,
			attachmentDao:    attachmentDao,
			contentDao:       contentDao,
			digDao:           digDao,
			forwardSourceDao: forwardSourceDao,
			paidUserDao:      paidUserDao,
			paymentDao:       paymentDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleDao *dao.ArticleDao,
		bookDao *dao.BookDao,
		chapterDao *dao.ChapterDao,
	) (*BookService, error) {
		obj := &BookService{
			articleDao: articleDao,
			bookDao:    bookDao,
			chapterDao: chapterDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		articleDao *dao.ArticleDao,
		chapterDao *dao.ChapterDao,
	) (*ChapterService, error) {
		obj := &ChapterService{
			articleDao: articleDao,
			chapterDao: chapterDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
