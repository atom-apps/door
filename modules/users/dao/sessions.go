package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/users/dto"

	"gorm.io/gen/field"
)

// @provider
type SessionDao struct {
	query *query.Query
}

func (dao *SessionDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *SessionDao) Context(ctx context.Context) query.ISessionDo {
	return dao.query.Session.WithContext(ctx)
}

func (dao *SessionDao) decorateSortQueryFilter(query query.ISessionDo, sortFilter *common.SortQueryFilter) query.ISessionDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Session.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Session.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *SessionDao) decorateQueryFilter(query query.ISessionDo, queryFilter *dto.SessionListQueryFilter) query.ISessionDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.Session.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.SessionID != nil {
		query = query.Where(dao.query.Session.SessionID.Eq(*queryFilter.SessionID))
	}

	return query
}

func (dao *SessionDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Session.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *SessionDao) Update(ctx context.Context, model *models.Session) error {
	_, err := dao.Context(ctx).Where(dao.query.Session.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *SessionDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Session.ID.Eq(id)).Delete()
	return err
}

func (dao *SessionDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Session.ID.Eq(id)).Delete()
	return err
}

func (dao *SessionDao) Create(ctx context.Context, model *models.Session) error {
	return dao.Context(ctx).Create(model)
}

func (dao *SessionDao) GetByID(ctx context.Context, id uint64) (*models.Session, error) {
	return dao.Context(ctx).Where(dao.query.Session.ID.Eq(id)).First()
}

func (dao *SessionDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Session, error) {
	return dao.Context(ctx).Where(dao.query.Session.ID.In(ids...)).Find()
}

func (dao *SessionDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.SessionListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Session, int64, error) {
	query := dao.query.Session
	sessionQuery := query.WithContext(ctx)
	sessionQuery = dao.decorateQueryFilter(sessionQuery, queryFilter)
	sessionQuery = dao.decorateSortQueryFilter(sessionQuery, sortFilter)
	return sessionQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *SessionDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.SessionListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Session, error) {
	query := dao.query.Session
	sessionQuery := query.WithContext(ctx)
	sessionQuery = dao.decorateQueryFilter(sessionQuery, queryFilter)
	sessionQuery = dao.decorateSortQueryFilter(sessionQuery, sortFilter)
	return sessionQuery.Find()
}

func (dao *SessionDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.SessionListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Session, error) {
	query := dao.query.Session
	sessionQuery := query.WithContext(ctx)
	sessionQuery = dao.decorateQueryFilter(sessionQuery, queryFilter)
	sessionQuery = dao.decorateSortQueryFilter(sessionQuery, sortFilter)
	return sessionQuery.First()
}

// GetBySessionID
func (dao *SessionDao) GetBySessionID(ctx context.Context, sessionID string) (*models.Session, error) {
	query := dao.query.Session
	return dao.Context(ctx).Where(query.SessionID.Eq(sessionID)).First()
}
