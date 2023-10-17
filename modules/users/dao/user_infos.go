package dao

import (
	"context"

	"github.com/atom-apps/door/common/ds"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/users/dto"

	"gorm.io/gen/field"
)

// @provider
type UserInfoDao struct {
	query *query.Query
}

func (dao *UserInfoDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *UserInfoDao) Context(ctx context.Context) query.IUserInfoDo {
	return dao.query.UserInfo.WithContext(ctx)
}

func (dao *UserInfoDao) decorateSortQueryFilter(query query.IUserInfoDo, sortFilter *ds.SortQueryFilter) query.IUserInfoDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.UserInfo.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.UserInfo.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *UserInfoDao) decorateQueryFilter(query query.IUserInfoDo, queryFilter *dto.UserInfoListQueryFilter) query.IUserInfoDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.UserInfo.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.Affiliation != nil {
		query = query.Where(dao.query.UserInfo.Affiliation.Eq(*queryFilter.Affiliation))
	}
	if queryFilter.Title != nil {
		query = query.Where(dao.query.UserInfo.Title.Eq(*queryFilter.Title))
	}
	if queryFilter.IDCardType != nil {
		query = query.Where(dao.query.UserInfo.IDCardType.Eq(*queryFilter.IDCardType))
	}
	if queryFilter.IDCard != nil {
		query = query.Where(dao.query.UserInfo.IDCard.Eq(*queryFilter.IDCard))
	}
	if queryFilter.Biography != nil {
		query = query.Where(dao.query.UserInfo.Biography.Eq(*queryFilter.Biography))
	}
	if queryFilter.Tag != nil {
		query = query.Where(dao.query.UserInfo.Tag.Eq(*queryFilter.Tag))
	}
	if queryFilter.Language != nil {
		query = query.Where(dao.query.UserInfo.Language.Eq(*queryFilter.Language))
	}
	if queryFilter.Gender != nil {
		query = query.Where(dao.query.UserInfo.Gender.Eq(*queryFilter.Gender))
	}
	if queryFilter.Birthday != nil {
		query = query.Where(dao.query.UserInfo.Birthday.Eq(*queryFilter.Birthday))
	}
	if queryFilter.Education != nil {
		query = query.Where(dao.query.UserInfo.Education.Eq(*queryFilter.Education))
	}
	if queryFilter.RealName != nil {
		query = query.Where(dao.query.UserInfo.RealName.Eq(*queryFilter.RealName))
	}

	return query
}

func (dao *UserInfoDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.UserInfo.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *UserInfoDao) Update(ctx context.Context, model *models.UserInfo) error {
	_, err := dao.Context(ctx).Where(dao.query.UserInfo.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *UserInfoDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.UserInfo.ID.Eq(id)).Delete()
	return err
}

func (dao *UserInfoDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.UserInfo.ID.Eq(id)).Delete()
	return err
}

func (dao *UserInfoDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.UserInfo.ID.Eq(id)).UpdateSimple(dao.query.UserInfo.DeletedAt.Null())
	return err
}

func (dao *UserInfoDao) Create(ctx context.Context, model *models.UserInfo) error {
	return dao.Context(ctx).Create(model)
}

func (dao *UserInfoDao) GetByID(ctx context.Context, id uint64) (*models.UserInfo, error) {
	return dao.Context(ctx).Where(dao.query.UserInfo.ID.Eq(id)).First()
}

func (dao *UserInfoDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.UserInfo, error) {
	return dao.Context(ctx).Where(dao.query.UserInfo.ID.In(ids...)).Find()
}

func (dao *UserInfoDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserInfoListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.UserInfo, int64, error) {
	query := dao.query.UserInfo
	userInfoQuery := query.WithContext(ctx)
	userInfoQuery = dao.decorateQueryFilter(userInfoQuery, queryFilter)
	userInfoQuery = dao.decorateSortQueryFilter(userInfoQuery, sortFilter)
	return userInfoQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *UserInfoDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserInfoListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.UserInfo, error) {
	query := dao.query.UserInfo
	userInfoQuery := query.WithContext(ctx)
	userInfoQuery = dao.decorateQueryFilter(userInfoQuery, queryFilter)
	userInfoQuery = dao.decorateSortQueryFilter(userInfoQuery, sortFilter)
	return userInfoQuery.Find()
}

func (dao *UserInfoDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserInfoListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.UserInfo, error) {
	query := dao.query.UserInfo
	userInfoQuery := query.WithContext(ctx)
	userInfoQuery = dao.decorateQueryFilter(userInfoQuery, queryFilter)
	userInfoQuery = dao.decorateSortQueryFilter(userInfoQuery, sortFilter)
	return userInfoQuery.First()
}
