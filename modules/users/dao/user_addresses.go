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
type UserAddressDao struct {
	query *query.Query
}

func (dao *UserAddressDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *UserAddressDao) Context(ctx context.Context) query.IUserAddressDo {
	return dao.query.UserAddress.WithContext(ctx)
}

func (dao *UserAddressDao) decorateSortQueryFilter(query query.IUserAddressDo, sortFilter *ds.SortQueryFilter) query.IUserAddressDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.UserAddress.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.UserAddress.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *UserAddressDao) decorateQueryFilter(query query.IUserAddressDo, queryFilter *dto.UserAddressListQueryFilter) query.IUserAddressDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.UserAddress.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.Code != nil {
		query = query.Where(dao.query.UserAddress.Code.Eq(*queryFilter.Code))
	}
	if queryFilter.Town != nil {
		query = query.Where(dao.query.UserAddress.Town.Eq(*queryFilter.Town))
	}
	if queryFilter.Detail != nil {
		query = query.Where(dao.query.UserAddress.Detail.Eq(*queryFilter.Detail))
	}
	if queryFilter.Name != nil {
		query = query.Where(dao.query.UserAddress.Name.Eq(*queryFilter.Name))
	}
	if queryFilter.Phone != nil {
		query = query.Where(dao.query.UserAddress.Phone.Eq(*queryFilter.Phone))
	}
	if queryFilter.ZipCode != nil {
		query = query.Where(dao.query.UserAddress.ZipCode.Eq(*queryFilter.ZipCode))
	}

	return query
}

func (dao *UserAddressDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.UserAddress.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *UserAddressDao) Update(ctx context.Context, model *models.UserAddress) error {
	_, err := dao.Context(ctx).Where(dao.query.UserAddress.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *UserAddressDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.UserAddress.ID.Eq(id)).Delete()
	return err
}

func (dao *UserAddressDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.UserAddress.ID.Eq(id)).Delete()
	return err
}

func (dao *UserAddressDao) Restore(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.UserAddress.ID.Eq(id)).UpdateSimple(dao.query.UserAddress.DeletedAt.Null())
	return err
}

func (dao *UserAddressDao) Create(ctx context.Context, model *models.UserAddress) error {
	return dao.Context(ctx).Create(model)
}

func (dao *UserAddressDao) GetByID(ctx context.Context, id uint64) (*models.UserAddress, error) {
	return dao.Context(ctx).Where(dao.query.UserAddress.ID.Eq(id)).First()
}

func (dao *UserAddressDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.UserAddress, error) {
	return dao.Context(ctx).Where(dao.query.UserAddress.ID.In(ids...)).Find()
}

func (dao *UserAddressDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserAddressListQueryFilter,
	pageFilter *ds.PageQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.UserAddress, int64, error) {
	query := dao.query.UserAddress
	userAddressQuery := query.WithContext(ctx)
	userAddressQuery = dao.decorateQueryFilter(userAddressQuery, queryFilter)
	userAddressQuery = dao.decorateSortQueryFilter(userAddressQuery, sortFilter)
	return userAddressQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *UserAddressDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserAddressListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) ([]*models.UserAddress, error) {
	query := dao.query.UserAddress
	userAddressQuery := query.WithContext(ctx)
	userAddressQuery = dao.decorateQueryFilter(userAddressQuery, queryFilter)
	userAddressQuery = dao.decorateSortQueryFilter(userAddressQuery, sortFilter)
	return userAddressQuery.Find()
}

func (dao *UserAddressDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserAddressListQueryFilter,
	sortFilter *ds.SortQueryFilter,
) (*models.UserAddress, error) {
	query := dao.query.UserAddress
	userAddressQuery := query.WithContext(ctx)
	userAddressQuery = dao.decorateQueryFilter(userAddressQuery, queryFilter)
	userAddressQuery = dao.decorateSortQueryFilter(userAddressQuery, sortFilter)
	return userAddressQuery.First()
}

func (dao *UserAddressDao) FindByUserID(ctx context.Context, userID uint64) ([]*models.UserAddress, error) {
	table, query := dao.query.UserAddress, dao.Context(ctx)
	return query.Where(table.UserID.Eq(userID)).Find()
}

func (dao *UserAddressDao) GetUserDefault(ctx context.Context, userID uint64) (*models.UserAddress, error) {
	table, query := dao.query.UserAddress, dao.Context(ctx)
	return query.Where(table.IsDefault.Is(true)).First()
}
