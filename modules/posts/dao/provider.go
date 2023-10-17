package dao

import (
	"github.com/atom-apps/door/database/query"
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ArticleAttachmentDao, error) {
		obj := &ArticleAttachmentDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ArticleContentDao, error) {
		obj := &ArticleContentDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ArticleDigDao, error) {
		obj := &ArticleDigDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ArticleForwardSourceDao, error) {
		obj := &ArticleForwardSourceDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ArticlePaidUserDao, error) {
		obj := &ArticlePaidUserDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ArticlePaymentDao, error) {
		obj := &ArticlePaymentDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ArticleTagDao, error) {
		obj := &ArticleTagDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
		tagDao *ArticleTagDao,
	) (*ArticleDao, error) {
		obj := &ArticleDao{
			query:  query,
			tagDao: tagDao,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*BookDao, error) {
		obj := &BookDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	if err := container.Container.Provide(func(
		query *query.Query,
	) (*ChapterDao, error) {
		obj := &ChapterDao{
			query: query,
		}
		return obj, nil
	}); err != nil {
		return err
	}

	return nil
}
