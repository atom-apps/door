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

	"github.com/atom-apps/door/database/database/models"
)

func newRole(db *gorm.DB, opts ...gen.DOOption) role {
	_role := role{}

	_role.roleDo.UseDB(db, opts...)
	_role.roleDo.UseModel(&models.Role{})

	tableName := _role.roleDo.TableName()
	_role.ALL = field.NewAsterisk(tableName)
	_role.ID = field.NewInt64(tableName, "id")
	_role.Name = field.NewString(tableName, "name")
	_role.Slug = field.NewString(tableName, "slug")
	_role.Description = field.NewString(tableName, "description")
	_role.ParentID = field.NewInt64(tableName, "parent_id")

	_role.fillFieldMap()

	return _role
}

type role struct {
	roleDo roleDo

	ALL         field.Asterisk
	ID          field.Int64
	Name        field.String
	Slug        field.String
	Description field.String
	ParentID    field.Int64

	fieldMap map[string]field.Expr
}

func (r role) Table(newTableName string) *role {
	r.roleDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r role) As(alias string) *role {
	r.roleDo.DO = *(r.roleDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *role) updateTableName(table string) *role {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.Name = field.NewString(table, "name")
	r.Slug = field.NewString(table, "slug")
	r.Description = field.NewString(table, "description")
	r.ParentID = field.NewInt64(table, "parent_id")

	r.fillFieldMap()

	return r
}

func (r *role) WithContext(ctx context.Context) IRoleDo { return r.roleDo.WithContext(ctx) }

func (r role) TableName() string { return r.roleDo.TableName() }

func (r role) Alias() string { return r.roleDo.Alias() }

func (r role) Columns(cols ...field.Expr) gen.Columns { return r.roleDo.Columns(cols...) }

func (r *role) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *role) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.ID
	r.fieldMap["name"] = r.Name
	r.fieldMap["slug"] = r.Slug
	r.fieldMap["description"] = r.Description
	r.fieldMap["parent_id"] = r.ParentID
}

func (r role) clone(db *gorm.DB) role {
	r.roleDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r role) replaceDB(db *gorm.DB) role {
	r.roleDo.ReplaceDB(db)
	return r
}

type roleDo struct{ gen.DO }

type IRoleDo interface {
	gen.SubQuery
	Debug() IRoleDo
	WithContext(ctx context.Context) IRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRoleDo
	WriteDB() IRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRoleDo
	Not(conds ...gen.Condition) IRoleDo
	Or(conds ...gen.Condition) IRoleDo
	Select(conds ...field.Expr) IRoleDo
	Where(conds ...gen.Condition) IRoleDo
	Order(conds ...field.Expr) IRoleDo
	Distinct(cols ...field.Expr) IRoleDo
	Omit(cols ...field.Expr) IRoleDo
	Join(table schema.Tabler, on ...field.Expr) IRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRoleDo
	Group(cols ...field.Expr) IRoleDo
	Having(conds ...gen.Condition) IRoleDo
	Limit(limit int) IRoleDo
	Offset(offset int) IRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleDo
	Unscoped() IRoleDo
	Create(values ...*models.Role) error
	CreateInBatches(values []*models.Role, batchSize int) error
	Save(values ...*models.Role) error
	First() (*models.Role, error)
	Take() (*models.Role, error)
	Last() (*models.Role, error)
	Find() ([]*models.Role, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Role, err error)
	FindInBatches(result *[]*models.Role, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Role) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRoleDo
	Assign(attrs ...field.AssignExpr) IRoleDo
	Joins(fields ...field.RelationField) IRoleDo
	Preload(fields ...field.RelationField) IRoleDo
	FirstOrInit() (*models.Role, error)
	FirstOrCreate() (*models.Role, error)
	FindByPage(offset int, limit int) (result []*models.Role, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r roleDo) Debug() IRoleDo {
	return r.withDO(r.DO.Debug())
}

func (r roleDo) WithContext(ctx context.Context) IRoleDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r roleDo) ReadDB() IRoleDo {
	return r.Clauses(dbresolver.Read)
}

func (r roleDo) WriteDB() IRoleDo {
	return r.Clauses(dbresolver.Write)
}

func (r roleDo) Session(config *gorm.Session) IRoleDo {
	return r.withDO(r.DO.Session(config))
}

func (r roleDo) Clauses(conds ...clause.Expression) IRoleDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r roleDo) Returning(value interface{}, columns ...string) IRoleDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r roleDo) Not(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r roleDo) Or(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r roleDo) Select(conds ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r roleDo) Where(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r roleDo) Order(conds ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r roleDo) Distinct(cols ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r roleDo) Omit(cols ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r roleDo) Join(table schema.Tabler, on ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r roleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRoleDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r roleDo) RightJoin(table schema.Tabler, on ...field.Expr) IRoleDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r roleDo) Group(cols ...field.Expr) IRoleDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r roleDo) Having(conds ...gen.Condition) IRoleDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r roleDo) Limit(limit int) IRoleDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r roleDo) Offset(offset int) IRoleDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r roleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRoleDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r roleDo) Unscoped() IRoleDo {
	return r.withDO(r.DO.Unscoped())
}

func (r roleDo) Create(values ...*models.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r roleDo) CreateInBatches(values []*models.Role, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r roleDo) Save(values ...*models.Role) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r roleDo) First() (*models.Role, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) Take() (*models.Role, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) Last() (*models.Role, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) Find() ([]*models.Role, error) {
	result, err := r.DO.Find()
	return result.([]*models.Role), err
}

func (r roleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Role, err error) {
	buf := make([]*models.Role, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r roleDo) FindInBatches(result *[]*models.Role, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r roleDo) Attrs(attrs ...field.AssignExpr) IRoleDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r roleDo) Assign(attrs ...field.AssignExpr) IRoleDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r roleDo) Joins(fields ...field.RelationField) IRoleDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r roleDo) Preload(fields ...field.RelationField) IRoleDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r roleDo) FirstOrInit() (*models.Role, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) FirstOrCreate() (*models.Role, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Role), nil
	}
}

func (r roleDo) FindByPage(offset int, limit int) (result []*models.Role, count int64, err error) {
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

func (r roleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r roleDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r roleDo) Delete(models ...*models.Role) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *roleDo) withDO(do gen.Dao) *roleDo {
	r.DO = *do.(*gen.DO)
	return r
}
