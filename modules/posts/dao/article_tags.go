package dao

import (
	"context"

	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/samber/lo"
	"gorm.io/gen/field"
)

// @provider
type ArticleTagDao struct {
	query *query.Query
}

func (dao *ArticleTagDao) Context(ctx context.Context) query.IArticleTagDo {
	return dao.query.ArticleTag.WithContext(ctx)
}

func (dao *ArticleTagDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleTag.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *ArticleTagDao) Update(ctx context.Context, model *models.ArticleTag) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleTag.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *ArticleTagDao) DeleteByArticleID(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.ArticleTag.ArticleID.Eq(id)).Delete()
	return err
}

func (dao *ArticleTagDao) Create(ctx context.Context, model *models.ArticleTag) error {
	return dao.Context(ctx).Create(model)
}

func (dao *ArticleTagDao) GetByID(ctx context.Context, id uint64) (*models.ArticleTag, error) {
	return dao.Context(ctx).Where(dao.query.ArticleTag.ID.Eq(id)).First()
}

func (dao *ArticleTagDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.ArticleTag, error) {
	return dao.Context(ctx).Where(dao.query.ArticleTag.ID.In(ids...)).Find()
}

func (dao *ArticleTagDao) GetByArticleIDs(ctx context.Context, ids []uint64) (map[uint64][]string, error) {
	aTag, tag := dao.query.ArticleTag, dao.query.Tag

	type articleTag struct {
		ArticleID uint64
		Name      string
	}
	var items []*articleTag
	err := dao.Context(ctx).
		Select(aTag.ArticleID, tag.Name).
		LeftJoin(tag, tag.ID.EqCol(aTag.TagID)).
		Where(aTag.ArticleID.In(ids...)).
		Scan(&items)
	if err != nil {
		return nil, err
	}

	maps := make(map[uint64][]string)
	lo.ForEach(items, func(item *articleTag, index int) {
		if _, ok := maps[item.ArticleID]; !ok {
			maps[item.ArticleID] = []string{}
		}
		maps[item.ArticleID] = append(maps[item.ArticleID], item.Name)
	})
	return maps, nil
}

func (dao *ArticleTagDao) GetByArticleID(ctx context.Context, id uint64) ([]string, error) {
	aTag, tag := dao.query.ArticleTag, dao.query.Tag

	var items []string
	err := dao.Context(ctx).
		Select(tag.Name).
		LeftJoin(tag, tag.ID.EqCol(aTag.TagID)).
		Where(aTag.ArticleID.Eq(id)).
		Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, err
}
