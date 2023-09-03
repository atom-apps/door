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
type TokenDao struct {
	query *query.Query
}

func (dao *TokenDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *TokenDao) Context(ctx context.Context) query.ITokenDo {
	return dao.query.Token.WithContext(ctx)
}

func (dao *TokenDao) decorateSortQueryFilter(query query.ITokenDo, sortFilter *common.SortQueryFilter) query.ITokenDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.Token.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.Token.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *TokenDao) decorateQueryFilter(query query.ITokenDo, queryFilter *dto.TokenListQueryFilter) query.ITokenDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.UserID != nil {
		query = query.Where(dao.query.Token.UserID.Eq(*queryFilter.UserID))
	}
	if queryFilter.AccessToken != nil {
		query = query.Where(dao.query.Token.AccessToken.Eq(*queryFilter.AccessToken))
	}
	if queryFilter.RefreshToken != nil {
		query = query.Where(dao.query.Token.RefreshToken.Eq(*queryFilter.RefreshToken))
	}
	if queryFilter.Scope != nil {
		query = query.Where(dao.query.Token.Scope.Eq(*queryFilter.Scope))
	}
	if queryFilter.TokenType != nil {
		query = query.Where(dao.query.Token.TokenType.Eq(*queryFilter.TokenType))
	}
	if queryFilter.CodeChallenge != nil {
		query = query.Where(dao.query.Token.CodeChallenge.Eq(*queryFilter.CodeChallenge))
	}
	if queryFilter.Code != nil {
		query = query.Where(dao.query.Token.Code.Eq(*queryFilter.Code))
	}
	if queryFilter.CodeExpireAt != nil {
		query = query.Where(dao.query.Token.CodeExpireAt.Eq(*queryFilter.CodeExpireAt))
	}
	if queryFilter.Used != nil {
		query = query.Where(dao.query.Token.Used.Is(*queryFilter.Used))
	}

	return query
}

func (dao *TokenDao) UpdateColumn(ctx context.Context, id uint64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.Token.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *TokenDao) Update(ctx context.Context, model *models.Token) error {
	_, err := dao.Context(ctx).Where(dao.query.Token.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *TokenDao) Delete(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Token.ID.Eq(id)).Delete()
	return err
}

func (dao *TokenDao) DeletePermanently(ctx context.Context, id uint64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.Token.ID.Eq(id)).Delete()
	return err
}

func (dao *TokenDao) Create(ctx context.Context, model *models.Token) error {
	return dao.Context(ctx).Create(model)
}

func (dao *TokenDao) GetByID(ctx context.Context, id uint64) (*models.Token, error) {
	return dao.Context(ctx).Where(dao.query.Token.ID.Eq(id)).First()
}

func (dao *TokenDao) GetByIDs(ctx context.Context, ids []uint64) ([]*models.Token, error) {
	return dao.Context(ctx).Where(dao.query.Token.ID.In(ids...)).Find()
}

func (dao *TokenDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TokenListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Token, int64, error) {
	query := dao.query.Token
	tokenQuery := query.WithContext(ctx)
	tokenQuery = dao.decorateQueryFilter(tokenQuery, queryFilter)
	tokenQuery = dao.decorateSortQueryFilter(tokenQuery, sortFilter)
	return tokenQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *TokenDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TokenListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.Token, error) {
	query := dao.query.Token
	tokenQuery := query.WithContext(ctx)
	tokenQuery = dao.decorateQueryFilter(tokenQuery, queryFilter)
	tokenQuery = dao.decorateSortQueryFilter(tokenQuery, sortFilter)
	return tokenQuery.Find()
}

func (dao *TokenDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.TokenListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.Token, error) {
	query := dao.query.Token
	tokenQuery := query.WithContext(ctx)
	tokenQuery = dao.decorateQueryFilter(tokenQuery, queryFilter)
	tokenQuery = dao.decorateSortQueryFilter(tokenQuery, sortFilter)
	return tokenQuery.First()
}

// GetByToken
func (dao *TokenDao) GetByToken(ctx context.Context, token string) (*models.Token, error) {
	table, query := dao.query.Token, dao.Context(ctx)
	return query.Where(table.AccessToken.Eq(token)).First()
}

func (dao *TokenDao) GetByUserID(ctx context.Context, userID uint64) (*models.Token, error) {
	table, query := dao.query.Token, dao.Context(ctx)
	return query.Where(table.UserID.Eq(userID)).First()
}

// DeleteBySessionID
func (dao *TokenDao) DeleteBySessionID(ctx context.Context, sessionID uint64) error {
	_, err := dao.Context(ctx).Where(dao.query.Token.SessionID.Eq(sessionID)).Delete()
	return err
}

// GetByRefreshToken
func (dao *TokenDao) GetByRefreshToken(ctx context.Context, refreshToken string) (*models.Token, error) {
	table, query := dao.query.Token, dao.Context(ctx)
	return query.Where(table.RefreshToken.Eq(refreshToken)).First()
}

// GetByCode
func (dao *TokenDao) GetByCode(ctx context.Context, code string) (*models.Token, error) {
	table, query := dao.query.Token, dao.Context(ctx)
	return query.Where(table.Code.Eq(code)).First()
}

// GetBySessionID
func (dao *TokenDao) GetBySessionID(ctx context.Context, sessionID uint64) (*models.Token, error) {
	table, query := dao.query.Token, dao.Context(ctx)
	return query.Where(table.SessionID.Eq(sessionID)).First()
}

// GetBySessionIDWithoutScope
func (dao *TokenDao) GetBySessionIDWithoutScope(ctx context.Context, sessionID uint64) ([]*models.Token, error) {
	table, query := dao.query.Token, dao.Context(ctx)
	return query.Where(table.SessionID.Eq(sessionID)).Find()
}
