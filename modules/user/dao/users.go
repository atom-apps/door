package dao

import (
	"context"

	"github.com/atom-apps/door/common"
	"github.com/atom-apps/door/common/consts"
	"github.com/atom-apps/door/database/models"
	"github.com/atom-apps/door/database/query"
	"github.com/atom-apps/door/modules/user/dto"

	"gorm.io/gen/field"
)

// @provider
type UserDao struct {
	query *query.Query
}

func (dao *UserDao) Transaction(f func() error) error {
	return dao.query.Transaction(func(tx *query.Query) error {
		return f()
	})
}

func (dao *UserDao) Context(ctx context.Context) query.IUserDo {
	return dao.query.User.WithContext(ctx)
}

func (dao *UserDao) decorateSortQueryFilter(query query.IUserDo, sortFilter *common.SortQueryFilter) query.IUserDo {
	if sortFilter == nil {
		return query
	}

	orderExprs := []field.Expr{}
	for _, v := range sortFilter.AscFields() {
		if expr, ok := dao.query.User.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr)
		}
	}
	for _, v := range sortFilter.DescFields() {
		if expr, ok := dao.query.User.GetFieldByName(v); ok {
			orderExprs = append(orderExprs, expr.Desc())
		}
	}
	return query.Order(orderExprs...)
}

func (dao *UserDao) decorateQueryFilter(query query.IUserDo, queryFilter *dto.UserListQueryFilter) query.IUserDo {
	if queryFilter == nil {
		return query
	}
	if queryFilter.UUID != nil {
		query = query.Where(dao.query.User.UUID.Eq(*queryFilter.UUID))
	}
	if queryFilter.Username != nil {
		query = query.Where(dao.query.User.Username.Eq(*queryFilter.Username))
	}
	if queryFilter.Password != nil {
		query = query.Where(dao.query.User.Password.Eq(*queryFilter.Password))
	}
	if queryFilter.Email != nil {
		query = query.Where(dao.query.User.Email.Eq(*queryFilter.Email))
	}
	if queryFilter.EmailVerified != nil {
		query = query.Where(dao.query.User.EmailVerified.Is(*queryFilter.EmailVerified))
	}
	if queryFilter.Phone != nil {
		query = query.Where(dao.query.User.Phone.Eq(*queryFilter.Phone))
	}
	if queryFilter.DisplayName != nil {
		query = query.Where(dao.query.User.DisplayName.Eq(*queryFilter.DisplayName))
	}
	if queryFilter.Avatar != nil {
		query = query.Where(dao.query.User.Avatar.Eq(*queryFilter.Avatar))
	}

	return query
}

func (dao *UserDao) UpdateColumn(ctx context.Context, id int64, field field.Expr, value interface{}) error {
	_, err := dao.Context(ctx).Where(dao.query.User.ID.Eq(id)).Update(field, value)
	return err
}

func (dao *UserDao) Update(ctx context.Context, model *models.User) error {
	_, err := dao.Context(ctx).Where(dao.query.User.ID.Eq(model.ID)).Updates(model)
	return err
}

func (dao *UserDao) Delete(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Where(dao.query.User.ID.Eq(id)).Delete()
	return err
}

func (dao *UserDao) DeletePermanently(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.User.ID.Eq(id)).Delete()
	return err
}

func (dao *UserDao) Restore(ctx context.Context, id int64) error {
	_, err := dao.Context(ctx).Unscoped().Where(dao.query.User.ID.Eq(id)).UpdateSimple(dao.query.User.DeletedAt.Null())
	return err
}

func (dao *UserDao) Create(ctx context.Context, model *models.User) error {
	return dao.Context(ctx).Create(model)
}

func (dao *UserDao) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return dao.Context(ctx).Where(dao.query.User.ID.Eq(id)).First()
}

func (dao *UserDao) GetByIDs(ctx context.Context, ids []int64) ([]*models.User, error) {
	return dao.Context(ctx).Where(dao.query.User.ID.In(ids...)).Find()
}

func (dao *UserDao) PageByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	pageFilter *common.PageQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, int64, error) {
	query := dao.query.User
	userQuery := query.WithContext(ctx)
	userQuery = dao.decorateQueryFilter(userQuery, queryFilter)
	userQuery = dao.decorateSortQueryFilter(userQuery, sortFilter)
	return userQuery.FindByPage(pageFilter.Offset(), pageFilter.Limit)
}

func (dao *UserDao) FindByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) ([]*models.User, error) {
	query := dao.query.User
	userQuery := query.WithContext(ctx)
	userQuery = dao.decorateQueryFilter(userQuery, queryFilter)
	userQuery = dao.decorateSortQueryFilter(userQuery, sortFilter)
	return userQuery.Find()
}

func (dao *UserDao) FirstByQueryFilter(
	ctx context.Context,
	queryFilter *dto.UserListQueryFilter,
	sortFilter *common.SortQueryFilter,
) (*models.User, error) {
	query := dao.query.User
	userQuery := query.WithContext(ctx)
	userQuery = dao.decorateQueryFilter(userQuery, queryFilter)
	userQuery = dao.decorateSortQueryFilter(userQuery, sortFilter)
	return userQuery.First()
}

// FindByEmail
func (dao *UserDao) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	query := dao.query.User
	return query.WithContext(ctx).Where(dao.query.User.Email.Eq(email)).First()
}

// FindByPhone
func (dao *UserDao) FindByPhone(ctx context.Context, phone string) (*models.User, error) {
	query := dao.query.User
	return query.WithContext(ctx).Where(dao.query.User.Phone.Eq(phone)).First()
}

// FindByUsername
func (dao *UserDao) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	query := dao.query.User
	return query.WithContext(ctx).Where(dao.query.User.Username.Eq(username)).First()
}

// GetByEmailOrPhone
func (dao *UserDao) GetByEmailOrPhone(ctx context.Context, emailOrPhone string) (*models.User, error) {
	user, query := dao.query.User, dao.Context(ctx)
	return query.Where(
		user.Status.Neq(consts.UserStatusBlocked),
		query.Where(user.Email.Eq(emailOrPhone)).Or(user.Phone.Eq(emailOrPhone)),
	).First()
}

// GetByEmailOrPhone
func (dao *UserDao) GetByUsernameOrEmailOrPhone(ctx context.Context, input string) (*models.User, error) {
	user, query := dao.query.User, dao.Context(ctx)
	return query.Where(
		user.Status.Neq(consts.UserStatusBlocked),
		query.Where(user.Username.Eq(input)).Or(user.Email.Eq(input)).Or(user.Phone.Eq(input)),
	).First()
}

// GetByPhone
func (dao *UserDao) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	user, query := dao.query.User, dao.Context(ctx)
	return query.Where(
		user.Status.Neq(consts.UserStatusBlocked),
		user.Phone.Eq(phone),
	).First()
}

// GetByEmail
func (dao *UserDao) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, query := dao.query.User, dao.Context(ctx)
	return query.Where(
		user.Status.Neq(consts.UserStatusBlocked),
		user.Email.Eq(email),
	).First()
}
