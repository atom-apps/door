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

func newDriver(db *gorm.DB, opts ...gen.DOOption) driver {
	_driver := driver{}

	_driver.driverDo.UseDB(db, opts...)
	_driver.driverDo.UseModel(&models.Driver{})

	tableName := _driver.driverDo.TableName()
	_driver.ALL = field.NewAsterisk(tableName)
	_driver.ID = field.NewUint64(tableName, "id")
	_driver.CreatedAt = field.NewTime(tableName, "created_at")
	_driver.UpdatedAt = field.NewTime(tableName, "updated_at")
	_driver.DeletedAt = field.NewField(tableName, "deleted_at")
	_driver.Type = field.NewField(tableName, "type")
	_driver.Name = field.NewString(tableName, "name")
	_driver.Endpoint = field.NewString(tableName, "endpoint")
	_driver.AccessKey = field.NewString(tableName, "access_key")
	_driver.AccessSecret = field.NewString(tableName, "access_secret")
	_driver.Bucket = field.NewString(tableName, "bucket")
	_driver.Options = field.NewField(tableName, "options")
	_driver.IsDefault = field.NewBool(tableName, "is_default")

	_driver.fillFieldMap()

	return _driver
}

type driver struct {
	driverDo driverDo

	ALL          field.Asterisk
	ID           field.Uint64 // ID
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间
	DeletedAt    field.Field  // 删除时间
	Type         field.Field  // 类型
	Name         field.String // 名称
	Endpoint     field.String // 地址
	AccessKey    field.String // AccessKey
	AccessSecret field.String // AccessSecret
	Bucket       field.String // Bucket
	Options      field.Field  // 配置
	IsDefault    field.Bool   // 默认

	fieldMap map[string]field.Expr
}

func (d driver) Table(newTableName string) *driver {
	d.driverDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d driver) As(alias string) *driver {
	d.driverDo.DO = *(d.driverDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *driver) updateTableName(table string) *driver {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewUint64(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Type = field.NewField(table, "type")
	d.Name = field.NewString(table, "name")
	d.Endpoint = field.NewString(table, "endpoint")
	d.AccessKey = field.NewString(table, "access_key")
	d.AccessSecret = field.NewString(table, "access_secret")
	d.Bucket = field.NewString(table, "bucket")
	d.Options = field.NewField(table, "options")
	d.IsDefault = field.NewBool(table, "is_default")

	d.fillFieldMap()

	return d
}

func (d *driver) WithContext(ctx context.Context) IDriverDo { return d.driverDo.WithContext(ctx) }

func (d driver) TableName() string { return d.driverDo.TableName() }

func (d driver) Alias() string { return d.driverDo.Alias() }

func (d driver) Columns(cols ...field.Expr) gen.Columns { return d.driverDo.Columns(cols...) }

func (d *driver) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *driver) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 12)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["type"] = d.Type
	d.fieldMap["name"] = d.Name
	d.fieldMap["endpoint"] = d.Endpoint
	d.fieldMap["access_key"] = d.AccessKey
	d.fieldMap["access_secret"] = d.AccessSecret
	d.fieldMap["bucket"] = d.Bucket
	d.fieldMap["options"] = d.Options
	d.fieldMap["is_default"] = d.IsDefault
}

func (d driver) clone(db *gorm.DB) driver {
	d.driverDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d driver) replaceDB(db *gorm.DB) driver {
	d.driverDo.ReplaceDB(db)
	return d
}

type driverDo struct{ gen.DO }

type IDriverDo interface {
	gen.SubQuery
	Debug() IDriverDo
	WithContext(ctx context.Context) IDriverDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDriverDo
	WriteDB() IDriverDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDriverDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDriverDo
	Not(conds ...gen.Condition) IDriverDo
	Or(conds ...gen.Condition) IDriverDo
	Select(conds ...field.Expr) IDriverDo
	Where(conds ...gen.Condition) IDriverDo
	Order(conds ...field.Expr) IDriverDo
	Distinct(cols ...field.Expr) IDriverDo
	Omit(cols ...field.Expr) IDriverDo
	Join(table schema.Tabler, on ...field.Expr) IDriverDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDriverDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDriverDo
	Group(cols ...field.Expr) IDriverDo
	Having(conds ...gen.Condition) IDriverDo
	Limit(limit int) IDriverDo
	Offset(offset int) IDriverDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDriverDo
	Unscoped() IDriverDo
	Create(values ...*models.Driver) error
	CreateInBatches(values []*models.Driver, batchSize int) error
	Save(values ...*models.Driver) error
	First() (*models.Driver, error)
	Take() (*models.Driver, error)
	Last() (*models.Driver, error)
	Find() ([]*models.Driver, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Driver, err error)
	FindInBatches(result *[]*models.Driver, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Driver) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDriverDo
	Assign(attrs ...field.AssignExpr) IDriverDo
	Joins(fields ...field.RelationField) IDriverDo
	Preload(fields ...field.RelationField) IDriverDo
	FirstOrInit() (*models.Driver, error)
	FirstOrCreate() (*models.Driver, error)
	FindByPage(offset int, limit int) (result []*models.Driver, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDriverDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d driverDo) Debug() IDriverDo {
	return d.withDO(d.DO.Debug())
}

func (d driverDo) WithContext(ctx context.Context) IDriverDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d driverDo) ReadDB() IDriverDo {
	return d.Clauses(dbresolver.Read)
}

func (d driverDo) WriteDB() IDriverDo {
	return d.Clauses(dbresolver.Write)
}

func (d driverDo) Session(config *gorm.Session) IDriverDo {
	return d.withDO(d.DO.Session(config))
}

func (d driverDo) Clauses(conds ...clause.Expression) IDriverDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d driverDo) Returning(value interface{}, columns ...string) IDriverDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d driverDo) Not(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d driverDo) Or(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d driverDo) Select(conds ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d driverDo) Where(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d driverDo) Order(conds ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d driverDo) Distinct(cols ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d driverDo) Omit(cols ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d driverDo) Join(table schema.Tabler, on ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d driverDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDriverDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d driverDo) RightJoin(table schema.Tabler, on ...field.Expr) IDriverDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d driverDo) Group(cols ...field.Expr) IDriverDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d driverDo) Having(conds ...gen.Condition) IDriverDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d driverDo) Limit(limit int) IDriverDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d driverDo) Offset(offset int) IDriverDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d driverDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDriverDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d driverDo) Unscoped() IDriverDo {
	return d.withDO(d.DO.Unscoped())
}

func (d driverDo) Create(values ...*models.Driver) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d driverDo) CreateInBatches(values []*models.Driver, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d driverDo) Save(values ...*models.Driver) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d driverDo) First() (*models.Driver, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Driver), nil
	}
}

func (d driverDo) Take() (*models.Driver, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Driver), nil
	}
}

func (d driverDo) Last() (*models.Driver, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Driver), nil
	}
}

func (d driverDo) Find() ([]*models.Driver, error) {
	result, err := d.DO.Find()
	return result.([]*models.Driver), err
}

func (d driverDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Driver, err error) {
	buf := make([]*models.Driver, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d driverDo) FindInBatches(result *[]*models.Driver, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d driverDo) Attrs(attrs ...field.AssignExpr) IDriverDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d driverDo) Assign(attrs ...field.AssignExpr) IDriverDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d driverDo) Joins(fields ...field.RelationField) IDriverDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d driverDo) Preload(fields ...field.RelationField) IDriverDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d driverDo) FirstOrInit() (*models.Driver, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Driver), nil
	}
}

func (d driverDo) FirstOrCreate() (*models.Driver, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Driver), nil
	}
}

func (d driverDo) FindByPage(offset int, limit int) (result []*models.Driver, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d driverDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d driverDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d driverDo) Delete(models ...*models.Driver) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *driverDo) withDO(do gen.Dao) *driverDo {
	d.DO = *do.(*gen.DO)
	return d
}
