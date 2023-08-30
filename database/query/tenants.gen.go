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

func newTenant(db *gorm.DB, opts ...gen.DOOption) tenant {
	_tenant := tenant{}

	_tenant.tenantDo.UseDB(db, opts...)
	_tenant.tenantDo.UseModel(&models.Tenant{})

	tableName := _tenant.tenantDo.TableName()
	_tenant.ALL = field.NewAsterisk(tableName)
	_tenant.ID = field.NewInt64(tableName, "id")
	_tenant.CreatedAt = field.NewTime(tableName, "created_at")
	_tenant.Name = field.NewString(tableName, "name")
	_tenant.Description = field.NewString(tableName, "description")
	_tenant.Meta = field.NewString(tableName, "meta")

	_tenant.fillFieldMap()

	return _tenant
}

type tenant struct {
	tenantDo tenantDo

	ALL         field.Asterisk
	ID          field.Int64  // ID
	CreatedAt   field.Time   // 创建时间
	Name        field.String // 名称
	Description field.String // 描述
	Meta        field.String // 元数据

	fieldMap map[string]field.Expr
}

func (t tenant) Table(newTableName string) *tenant {
	t.tenantDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tenant) As(alias string) *tenant {
	t.tenantDo.DO = *(t.tenantDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tenant) updateTableName(table string) *tenant {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.Name = field.NewString(table, "name")
	t.Description = field.NewString(table, "description")
	t.Meta = field.NewString(table, "meta")

	t.fillFieldMap()

	return t
}

func (t *tenant) WithContext(ctx context.Context) ITenantDo { return t.tenantDo.WithContext(ctx) }

func (t tenant) TableName() string { return t.tenantDo.TableName() }

func (t tenant) Alias() string { return t.tenantDo.Alias() }

func (t tenant) Columns(cols ...field.Expr) gen.Columns { return t.tenantDo.Columns(cols...) }

func (t *tenant) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tenant) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["id"] = t.ID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["name"] = t.Name
	t.fieldMap["description"] = t.Description
	t.fieldMap["meta"] = t.Meta
}

func (t tenant) clone(db *gorm.DB) tenant {
	t.tenantDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tenant) replaceDB(db *gorm.DB) tenant {
	t.tenantDo.ReplaceDB(db)
	return t
}

type tenantDo struct{ gen.DO }

type ITenantDo interface {
	gen.SubQuery
	Debug() ITenantDo
	WithContext(ctx context.Context) ITenantDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITenantDo
	WriteDB() ITenantDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITenantDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITenantDo
	Not(conds ...gen.Condition) ITenantDo
	Or(conds ...gen.Condition) ITenantDo
	Select(conds ...field.Expr) ITenantDo
	Where(conds ...gen.Condition) ITenantDo
	Order(conds ...field.Expr) ITenantDo
	Distinct(cols ...field.Expr) ITenantDo
	Omit(cols ...field.Expr) ITenantDo
	Join(table schema.Tabler, on ...field.Expr) ITenantDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITenantDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITenantDo
	Group(cols ...field.Expr) ITenantDo
	Having(conds ...gen.Condition) ITenantDo
	Limit(limit int) ITenantDo
	Offset(offset int) ITenantDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITenantDo
	Unscoped() ITenantDo
	Create(values ...*models.Tenant) error
	CreateInBatches(values []*models.Tenant, batchSize int) error
	Save(values ...*models.Tenant) error
	First() (*models.Tenant, error)
	Take() (*models.Tenant, error)
	Last() (*models.Tenant, error)
	Find() ([]*models.Tenant, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Tenant, err error)
	FindInBatches(result *[]*models.Tenant, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Tenant) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITenantDo
	Assign(attrs ...field.AssignExpr) ITenantDo
	Joins(fields ...field.RelationField) ITenantDo
	Preload(fields ...field.RelationField) ITenantDo
	FirstOrInit() (*models.Tenant, error)
	FirstOrCreate() (*models.Tenant, error)
	FindByPage(offset int, limit int) (result []*models.Tenant, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITenantDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tenantDo) Debug() ITenantDo {
	return t.withDO(t.DO.Debug())
}

func (t tenantDo) WithContext(ctx context.Context) ITenantDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tenantDo) ReadDB() ITenantDo {
	return t.Clauses(dbresolver.Read)
}

func (t tenantDo) WriteDB() ITenantDo {
	return t.Clauses(dbresolver.Write)
}

func (t tenantDo) Session(config *gorm.Session) ITenantDo {
	return t.withDO(t.DO.Session(config))
}

func (t tenantDo) Clauses(conds ...clause.Expression) ITenantDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tenantDo) Returning(value interface{}, columns ...string) ITenantDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tenantDo) Not(conds ...gen.Condition) ITenantDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tenantDo) Or(conds ...gen.Condition) ITenantDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tenantDo) Select(conds ...field.Expr) ITenantDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tenantDo) Where(conds ...gen.Condition) ITenantDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tenantDo) Order(conds ...field.Expr) ITenantDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tenantDo) Distinct(cols ...field.Expr) ITenantDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tenantDo) Omit(cols ...field.Expr) ITenantDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tenantDo) Join(table schema.Tabler, on ...field.Expr) ITenantDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tenantDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITenantDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tenantDo) RightJoin(table schema.Tabler, on ...field.Expr) ITenantDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tenantDo) Group(cols ...field.Expr) ITenantDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tenantDo) Having(conds ...gen.Condition) ITenantDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tenantDo) Limit(limit int) ITenantDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tenantDo) Offset(offset int) ITenantDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tenantDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITenantDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tenantDo) Unscoped() ITenantDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tenantDo) Create(values ...*models.Tenant) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tenantDo) CreateInBatches(values []*models.Tenant, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tenantDo) Save(values ...*models.Tenant) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tenantDo) First() (*models.Tenant, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Tenant), nil
	}
}

func (t tenantDo) Take() (*models.Tenant, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Tenant), nil
	}
}

func (t tenantDo) Last() (*models.Tenant, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Tenant), nil
	}
}

func (t tenantDo) Find() ([]*models.Tenant, error) {
	result, err := t.DO.Find()
	return result.([]*models.Tenant), err
}

func (t tenantDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Tenant, err error) {
	buf := make([]*models.Tenant, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tenantDo) FindInBatches(result *[]*models.Tenant, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tenantDo) Attrs(attrs ...field.AssignExpr) ITenantDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tenantDo) Assign(attrs ...field.AssignExpr) ITenantDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tenantDo) Joins(fields ...field.RelationField) ITenantDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tenantDo) Preload(fields ...field.RelationField) ITenantDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tenantDo) FirstOrInit() (*models.Tenant, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Tenant), nil
	}
}

func (t tenantDo) FirstOrCreate() (*models.Tenant, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Tenant), nil
	}
}

func (t tenantDo) FindByPage(offset int, limit int) (result []*models.Tenant, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tenantDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tenantDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tenantDo) Delete(models ...*models.Tenant) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tenantDo) withDO(do gen.Dao) *tenantDo {
	t.DO = *do.(*gen.DO)
	return t
}
