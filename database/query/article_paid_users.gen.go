// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/atom-apps/door/database/models"
)

func newArticlePaidUser(db *gorm.DB, opts ...gen.DOOption) articlePaidUser {
	_articlePaidUser := articlePaidUser{}

	_articlePaidUser.articlePaidUserDo.UseDB(db, opts...)
	_articlePaidUser.articlePaidUserDo.UseModel(&models.ArticlePaidUser{})

	tableName := _articlePaidUser.articlePaidUserDo.TableName()
	_articlePaidUser.ALL = field.NewAsterisk(tableName)
	_articlePaidUser.ID = field.NewUint64(tableName, "id")
	_articlePaidUser.CreatedAt = field.NewTime(tableName, "created_at")
	_articlePaidUser.TenantID = field.NewUint64(tableName, "tenant_id")
	_articlePaidUser.UserID = field.NewUint64(tableName, "user_id")
	_articlePaidUser.ArticleID = field.NewUint64(tableName, "article_id")
	_articlePaidUser.PriceType = field.NewInt64(tableName, "price_type")
	_articlePaidUser.Price = field.NewUint64(tableName, "price")
	_articlePaidUser.EndAt = field.NewUint64(tableName, "end_at")

	_articlePaidUser.fillFieldMap()

	return _articlePaidUser
}

type articlePaidUser struct {
	articlePaidUserDo articlePaidUserDo

	ALL       field.Asterisk
	ID        field.Uint64 // ID
	CreatedAt field.Time   // 创建时间
	TenantID  field.Uint64 // 租户ID
	UserID    field.Uint64 // 用户ID
	ArticleID field.Uint64 // 文章ID
	PriceType field.Int64  // 付费类型
	Price     field.Uint64 // 付费价格
	EndAt     field.Uint64 // 付费结束时间

	fieldMap map[string]field.Expr
}

func (a articlePaidUser) Table(newTableName string) *articlePaidUser {
	a.articlePaidUserDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articlePaidUser) As(alias string) *articlePaidUser {
	a.articlePaidUserDo.DO = *(a.articlePaidUserDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articlePaidUser) updateTableName(table string) *articlePaidUser {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.TenantID = field.NewUint64(table, "tenant_id")
	a.UserID = field.NewUint64(table, "user_id")
	a.ArticleID = field.NewUint64(table, "article_id")
	a.PriceType = field.NewInt64(table, "price_type")
	a.Price = field.NewUint64(table, "price")
	a.EndAt = field.NewUint64(table, "end_at")

	a.fillFieldMap()

	return a
}

func (a *articlePaidUser) WithContext(ctx context.Context) IArticlePaidUserDo {
	return a.articlePaidUserDo.WithContext(ctx)
}

func (a articlePaidUser) TableName() string { return a.articlePaidUserDo.TableName() }

func (a articlePaidUser) Alias() string { return a.articlePaidUserDo.Alias() }

func (a articlePaidUser) Columns(cols ...field.Expr) gen.Columns {
	return a.articlePaidUserDo.Columns(cols...)
}

func (a *articlePaidUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articlePaidUser) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["tenant_id"] = a.TenantID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["price_type"] = a.PriceType
	a.fieldMap["price"] = a.Price
	a.fieldMap["end_at"] = a.EndAt
}

func (a articlePaidUser) clone(db *gorm.DB) articlePaidUser {
	a.articlePaidUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articlePaidUser) replaceDB(db *gorm.DB) articlePaidUser {
	a.articlePaidUserDo.ReplaceDB(db)
	return a
}

type articlePaidUserDo struct{ gen.DO }

type IArticlePaidUserDo interface {
	gen.SubQuery
	Debug() IArticlePaidUserDo
	WithContext(ctx context.Context) IArticlePaidUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticlePaidUserDo
	WriteDB() IArticlePaidUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticlePaidUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticlePaidUserDo
	Not(conds ...gen.Condition) IArticlePaidUserDo
	Or(conds ...gen.Condition) IArticlePaidUserDo
	Select(conds ...field.Expr) IArticlePaidUserDo
	Where(conds ...gen.Condition) IArticlePaidUserDo
	Order(conds ...field.Expr) IArticlePaidUserDo
	Distinct(cols ...field.Expr) IArticlePaidUserDo
	Omit(cols ...field.Expr) IArticlePaidUserDo
	Join(table schema.Tabler, on ...field.Expr) IArticlePaidUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticlePaidUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticlePaidUserDo
	Group(cols ...field.Expr) IArticlePaidUserDo
	Having(conds ...gen.Condition) IArticlePaidUserDo
	Limit(limit int) IArticlePaidUserDo
	Offset(offset int) IArticlePaidUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticlePaidUserDo
	Unscoped() IArticlePaidUserDo
	Create(values ...*models.ArticlePaidUser) error
	CreateInBatches(values []*models.ArticlePaidUser, batchSize int) error
	Save(values ...*models.ArticlePaidUser) error
	First() (*models.ArticlePaidUser, error)
	Take() (*models.ArticlePaidUser, error)
	Last() (*models.ArticlePaidUser, error)
	Find() ([]*models.ArticlePaidUser, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticlePaidUser, err error)
	FindInBatches(result *[]*models.ArticlePaidUser, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.ArticlePaidUser) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticlePaidUserDo
	Assign(attrs ...field.AssignExpr) IArticlePaidUserDo
	Joins(fields ...field.RelationField) IArticlePaidUserDo
	Preload(fields ...field.RelationField) IArticlePaidUserDo
	FirstOrInit() (*models.ArticlePaidUser, error)
	FirstOrCreate() (*models.ArticlePaidUser, error)
	FindByPage(offset int, limit int) (result []*models.ArticlePaidUser, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticlePaidUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articlePaidUserDo) Debug() IArticlePaidUserDo {
	return a.withDO(a.DO.Debug())
}

func (a articlePaidUserDo) WithContext(ctx context.Context) IArticlePaidUserDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articlePaidUserDo) ReadDB() IArticlePaidUserDo {
	return a.Clauses(dbresolver.Read)
}

func (a articlePaidUserDo) WriteDB() IArticlePaidUserDo {
	return a.Clauses(dbresolver.Write)
}

func (a articlePaidUserDo) Session(config *gorm.Session) IArticlePaidUserDo {
	return a.withDO(a.DO.Session(config))
}

func (a articlePaidUserDo) Clauses(conds ...clause.Expression) IArticlePaidUserDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articlePaidUserDo) Returning(value interface{}, columns ...string) IArticlePaidUserDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articlePaidUserDo) Not(conds ...gen.Condition) IArticlePaidUserDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articlePaidUserDo) Or(conds ...gen.Condition) IArticlePaidUserDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articlePaidUserDo) Select(conds ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articlePaidUserDo) Where(conds ...gen.Condition) IArticlePaidUserDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articlePaidUserDo) Order(conds ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articlePaidUserDo) Distinct(cols ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articlePaidUserDo) Omit(cols ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articlePaidUserDo) Join(table schema.Tabler, on ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articlePaidUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articlePaidUserDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articlePaidUserDo) Group(cols ...field.Expr) IArticlePaidUserDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articlePaidUserDo) Having(conds ...gen.Condition) IArticlePaidUserDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articlePaidUserDo) Limit(limit int) IArticlePaidUserDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articlePaidUserDo) Offset(offset int) IArticlePaidUserDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articlePaidUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticlePaidUserDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articlePaidUserDo) Unscoped() IArticlePaidUserDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articlePaidUserDo) Create(values ...*models.ArticlePaidUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articlePaidUserDo) CreateInBatches(values []*models.ArticlePaidUser, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articlePaidUserDo) Save(values ...*models.ArticlePaidUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articlePaidUserDo) First() (*models.ArticlePaidUser, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePaidUser), nil
	}
}

func (a articlePaidUserDo) Take() (*models.ArticlePaidUser, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePaidUser), nil
	}
}

func (a articlePaidUserDo) Last() (*models.ArticlePaidUser, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePaidUser), nil
	}
}

func (a articlePaidUserDo) Find() ([]*models.ArticlePaidUser, error) {
	result, err := a.DO.Find()
	return result.([]*models.ArticlePaidUser), err
}

func (a articlePaidUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.ArticlePaidUser, err error) {
	buf := make([]*models.ArticlePaidUser, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articlePaidUserDo) FindInBatches(result *[]*models.ArticlePaidUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articlePaidUserDo) Attrs(attrs ...field.AssignExpr) IArticlePaidUserDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articlePaidUserDo) Assign(attrs ...field.AssignExpr) IArticlePaidUserDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articlePaidUserDo) Joins(fields ...field.RelationField) IArticlePaidUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articlePaidUserDo) Preload(fields ...field.RelationField) IArticlePaidUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articlePaidUserDo) FirstOrInit() (*models.ArticlePaidUser, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePaidUser), nil
	}
}

func (a articlePaidUserDo) FirstOrCreate() (*models.ArticlePaidUser, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.ArticlePaidUser), nil
	}
}

func (a articlePaidUserDo) FindByPage(offset int, limit int) (result []*models.ArticlePaidUser, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articlePaidUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articlePaidUserDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articlePaidUserDo) Delete(models ...*models.ArticlePaidUser) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articlePaidUserDo) withDO(do gen.Dao) *articlePaidUserDo {
	a.DO = *do.(*gen.DO)
	return a
}
