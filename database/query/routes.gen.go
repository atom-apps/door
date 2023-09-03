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

func newRoute(db *gorm.DB, opts ...gen.DOOption) route {
	_route := route{}

	_route.routeDo.UseDB(db, opts...)
	_route.routeDo.UseModel(&models.Route{})

	tableName := _route.routeDo.TableName()
	_route.ALL = field.NewAsterisk(tableName)
	_route.ID = field.NewUint64(tableName, "id")
	_route.Type = field.NewField(tableName, "type")
	_route.ParentID = field.NewUint64(tableName, "parent_id")
	_route.Name = field.NewString(tableName, "name")
	_route.Path = field.NewString(tableName, "path")
	_route.Metadata = field.NewField(tableName, "metadata")
	_route.Order = field.NewUint64(tableName, "order")

	_route.fillFieldMap()

	return _route
}

type route struct {
	routeDo routeDo

	ALL      field.Asterisk
	ID       field.Uint64
	Type     field.Field
	ParentID field.Uint64
	Name     field.String
	Path     field.String
	Metadata field.Field
	Order    field.Uint64

	fieldMap map[string]field.Expr
}

func (r route) Table(newTableName string) *route {
	r.routeDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r route) As(alias string) *route {
	r.routeDo.DO = *(r.routeDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *route) updateTableName(table string) *route {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint64(table, "id")
	r.Type = field.NewField(table, "type")
	r.ParentID = field.NewUint64(table, "parent_id")
	r.Name = field.NewString(table, "name")
	r.Path = field.NewString(table, "path")
	r.Metadata = field.NewField(table, "metadata")
	r.Order = field.NewUint64(table, "order")

	r.fillFieldMap()

	return r
}

func (r *route) WithContext(ctx context.Context) IRouteDo { return r.routeDo.WithContext(ctx) }

func (r route) TableName() string { return r.routeDo.TableName() }

func (r route) Alias() string { return r.routeDo.Alias() }

func (r route) Columns(cols ...field.Expr) gen.Columns { return r.routeDo.Columns(cols...) }

func (r *route) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *route) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 7)
	r.fieldMap["id"] = r.ID
	r.fieldMap["type"] = r.Type
	r.fieldMap["parent_id"] = r.ParentID
	r.fieldMap["name"] = r.Name
	r.fieldMap["path"] = r.Path
	r.fieldMap["metadata"] = r.Metadata
	r.fieldMap["order"] = r.Order
}

func (r route) clone(db *gorm.DB) route {
	r.routeDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r route) replaceDB(db *gorm.DB) route {
	r.routeDo.ReplaceDB(db)
	return r
}

type routeDo struct{ gen.DO }

type IRouteDo interface {
	gen.SubQuery
	Debug() IRouteDo
	WithContext(ctx context.Context) IRouteDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRouteDo
	WriteDB() IRouteDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRouteDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRouteDo
	Not(conds ...gen.Condition) IRouteDo
	Or(conds ...gen.Condition) IRouteDo
	Select(conds ...field.Expr) IRouteDo
	Where(conds ...gen.Condition) IRouteDo
	Order(conds ...field.Expr) IRouteDo
	Distinct(cols ...field.Expr) IRouteDo
	Omit(cols ...field.Expr) IRouteDo
	Join(table schema.Tabler, on ...field.Expr) IRouteDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRouteDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRouteDo
	Group(cols ...field.Expr) IRouteDo
	Having(conds ...gen.Condition) IRouteDo
	Limit(limit int) IRouteDo
	Offset(offset int) IRouteDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRouteDo
	Unscoped() IRouteDo
	Create(values ...*models.Route) error
	CreateInBatches(values []*models.Route, batchSize int) error
	Save(values ...*models.Route) error
	First() (*models.Route, error)
	Take() (*models.Route, error)
	Last() (*models.Route, error)
	Find() ([]*models.Route, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Route, err error)
	FindInBatches(result *[]*models.Route, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Route) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRouteDo
	Assign(attrs ...field.AssignExpr) IRouteDo
	Joins(fields ...field.RelationField) IRouteDo
	Preload(fields ...field.RelationField) IRouteDo
	FirstOrInit() (*models.Route, error)
	FirstOrCreate() (*models.Route, error)
	FindByPage(offset int, limit int) (result []*models.Route, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRouteDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r routeDo) Debug() IRouteDo {
	return r.withDO(r.DO.Debug())
}

func (r routeDo) WithContext(ctx context.Context) IRouteDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r routeDo) ReadDB() IRouteDo {
	return r.Clauses(dbresolver.Read)
}

func (r routeDo) WriteDB() IRouteDo {
	return r.Clauses(dbresolver.Write)
}

func (r routeDo) Session(config *gorm.Session) IRouteDo {
	return r.withDO(r.DO.Session(config))
}

func (r routeDo) Clauses(conds ...clause.Expression) IRouteDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r routeDo) Returning(value interface{}, columns ...string) IRouteDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r routeDo) Not(conds ...gen.Condition) IRouteDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r routeDo) Or(conds ...gen.Condition) IRouteDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r routeDo) Select(conds ...field.Expr) IRouteDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r routeDo) Where(conds ...gen.Condition) IRouteDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r routeDo) Order(conds ...field.Expr) IRouteDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r routeDo) Distinct(cols ...field.Expr) IRouteDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r routeDo) Omit(cols ...field.Expr) IRouteDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r routeDo) Join(table schema.Tabler, on ...field.Expr) IRouteDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r routeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRouteDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r routeDo) RightJoin(table schema.Tabler, on ...field.Expr) IRouteDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r routeDo) Group(cols ...field.Expr) IRouteDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r routeDo) Having(conds ...gen.Condition) IRouteDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r routeDo) Limit(limit int) IRouteDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r routeDo) Offset(offset int) IRouteDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r routeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRouteDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r routeDo) Unscoped() IRouteDo {
	return r.withDO(r.DO.Unscoped())
}

func (r routeDo) Create(values ...*models.Route) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r routeDo) CreateInBatches(values []*models.Route, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r routeDo) Save(values ...*models.Route) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r routeDo) First() (*models.Route, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Route), nil
	}
}

func (r routeDo) Take() (*models.Route, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Route), nil
	}
}

func (r routeDo) Last() (*models.Route, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Route), nil
	}
}

func (r routeDo) Find() ([]*models.Route, error) {
	result, err := r.DO.Find()
	return result.([]*models.Route), err
}

func (r routeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Route, err error) {
	buf := make([]*models.Route, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r routeDo) FindInBatches(result *[]*models.Route, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r routeDo) Attrs(attrs ...field.AssignExpr) IRouteDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r routeDo) Assign(attrs ...field.AssignExpr) IRouteDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r routeDo) Joins(fields ...field.RelationField) IRouteDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r routeDo) Preload(fields ...field.RelationField) IRouteDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r routeDo) FirstOrInit() (*models.Route, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Route), nil
	}
}

func (r routeDo) FirstOrCreate() (*models.Route, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Route), nil
	}
}

func (r routeDo) FindByPage(offset int, limit int) (result []*models.Route, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r routeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r routeDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r routeDo) Delete(models ...*models.Route) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *routeDo) withDO(do gen.Dao) *routeDo {
	r.DO = *do.(*gen.DO)
	return r
}
