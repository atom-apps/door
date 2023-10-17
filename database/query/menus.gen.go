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

func newMenu(db *gorm.DB, opts ...gen.DOOption) menu {
	_menu := menu{}

	_menu.menuDo.UseDB(db, opts...)
	_menu.menuDo.UseModel(&models.Menu{})

	tableName := _menu.menuDo.TableName()
	_menu.ALL = field.NewAsterisk(tableName)
	_menu.ID = field.NewUint64(tableName, "id")
	_menu.CreatedAt = field.NewTime(tableName, "created_at")
	_menu.Name = field.NewString(tableName, "name")
	_menu.Slug = field.NewString(tableName, "slug")
	_menu.GroupID = field.NewUint64(tableName, "group_id")
	_menu.ParentID = field.NewUint64(tableName, "parent_id")
	_menu.Metadata = field.NewField(tableName, "metadata")

	_menu.fillFieldMap()

	return _menu
}

type menu struct {
	menuDo menuDo

	ALL       field.Asterisk
	ID        field.Uint64 // ID
	CreatedAt field.Time   // 创建时间
	Name      field.String // 名称
	Slug      field.String // 别名
	GroupID   field.Uint64 // 组
	ParentID  field.Uint64 // 父ID
	Metadata  field.Field  // 元数据

	fieldMap map[string]field.Expr
}

func (m menu) Table(newTableName string) *menu {
	m.menuDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m menu) As(alias string) *menu {
	m.menuDo.DO = *(m.menuDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *menu) updateTableName(table string) *menu {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint64(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.Name = field.NewString(table, "name")
	m.Slug = field.NewString(table, "slug")
	m.GroupID = field.NewUint64(table, "group_id")
	m.ParentID = field.NewUint64(table, "parent_id")
	m.Metadata = field.NewField(table, "metadata")

	m.fillFieldMap()

	return m
}

func (m *menu) WithContext(ctx context.Context) IMenuDo { return m.menuDo.WithContext(ctx) }

func (m menu) TableName() string { return m.menuDo.TableName() }

func (m menu) Alias() string { return m.menuDo.Alias() }

func (m menu) Columns(cols ...field.Expr) gen.Columns { return m.menuDo.Columns(cols...) }

func (m *menu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *menu) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 7)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["name"] = m.Name
	m.fieldMap["slug"] = m.Slug
	m.fieldMap["group_id"] = m.GroupID
	m.fieldMap["parent_id"] = m.ParentID
	m.fieldMap["metadata"] = m.Metadata
}

func (m menu) clone(db *gorm.DB) menu {
	m.menuDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m menu) replaceDB(db *gorm.DB) menu {
	m.menuDo.ReplaceDB(db)
	return m
}

type menuDo struct{ gen.DO }

type IMenuDo interface {
	gen.SubQuery
	Debug() IMenuDo
	WithContext(ctx context.Context) IMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMenuDo
	WriteDB() IMenuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMenuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMenuDo
	Not(conds ...gen.Condition) IMenuDo
	Or(conds ...gen.Condition) IMenuDo
	Select(conds ...field.Expr) IMenuDo
	Where(conds ...gen.Condition) IMenuDo
	Order(conds ...field.Expr) IMenuDo
	Distinct(cols ...field.Expr) IMenuDo
	Omit(cols ...field.Expr) IMenuDo
	Join(table schema.Tabler, on ...field.Expr) IMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMenuDo
	Group(cols ...field.Expr) IMenuDo
	Having(conds ...gen.Condition) IMenuDo
	Limit(limit int) IMenuDo
	Offset(offset int) IMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMenuDo
	Unscoped() IMenuDo
	Create(values ...*models.Menu) error
	CreateInBatches(values []*models.Menu, batchSize int) error
	Save(values ...*models.Menu) error
	First() (*models.Menu, error)
	Take() (*models.Menu, error)
	Last() (*models.Menu, error)
	Find() ([]*models.Menu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Menu, err error)
	FindInBatches(result *[]*models.Menu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Menu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMenuDo
	Assign(attrs ...field.AssignExpr) IMenuDo
	Joins(fields ...field.RelationField) IMenuDo
	Preload(fields ...field.RelationField) IMenuDo
	FirstOrInit() (*models.Menu, error)
	FirstOrCreate() (*models.Menu, error)
	FindByPage(offset int, limit int) (result []*models.Menu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m menuDo) Debug() IMenuDo {
	return m.withDO(m.DO.Debug())
}

func (m menuDo) WithContext(ctx context.Context) IMenuDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m menuDo) ReadDB() IMenuDo {
	return m.Clauses(dbresolver.Read)
}

func (m menuDo) WriteDB() IMenuDo {
	return m.Clauses(dbresolver.Write)
}

func (m menuDo) Session(config *gorm.Session) IMenuDo {
	return m.withDO(m.DO.Session(config))
}

func (m menuDo) Clauses(conds ...clause.Expression) IMenuDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m menuDo) Returning(value interface{}, columns ...string) IMenuDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m menuDo) Not(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m menuDo) Or(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m menuDo) Select(conds ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m menuDo) Where(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m menuDo) Order(conds ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m menuDo) Distinct(cols ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m menuDo) Omit(cols ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m menuDo) Join(table schema.Tabler, on ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m menuDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMenuDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m menuDo) RightJoin(table schema.Tabler, on ...field.Expr) IMenuDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m menuDo) Group(cols ...field.Expr) IMenuDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m menuDo) Having(conds ...gen.Condition) IMenuDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m menuDo) Limit(limit int) IMenuDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m menuDo) Offset(offset int) IMenuDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m menuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMenuDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m menuDo) Unscoped() IMenuDo {
	return m.withDO(m.DO.Unscoped())
}

func (m menuDo) Create(values ...*models.Menu) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m menuDo) CreateInBatches(values []*models.Menu, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m menuDo) Save(values ...*models.Menu) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m menuDo) First() (*models.Menu, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Menu), nil
	}
}

func (m menuDo) Take() (*models.Menu, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Menu), nil
	}
}

func (m menuDo) Last() (*models.Menu, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Menu), nil
	}
}

func (m menuDo) Find() ([]*models.Menu, error) {
	result, err := m.DO.Find()
	return result.([]*models.Menu), err
}

func (m menuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Menu, err error) {
	buf := make([]*models.Menu, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m menuDo) FindInBatches(result *[]*models.Menu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m menuDo) Attrs(attrs ...field.AssignExpr) IMenuDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m menuDo) Assign(attrs ...field.AssignExpr) IMenuDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m menuDo) Joins(fields ...field.RelationField) IMenuDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m menuDo) Preload(fields ...field.RelationField) IMenuDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m menuDo) FirstOrInit() (*models.Menu, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Menu), nil
	}
}

func (m menuDo) FirstOrCreate() (*models.Menu, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Menu), nil
	}
}

func (m menuDo) FindByPage(offset int, limit int) (result []*models.Menu, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m menuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m menuDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m menuDo) Delete(models ...*models.Menu) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *menuDo) withDO(do gen.Dao) *menuDo {
	m.DO = *do.(*gen.DO)
	return m
}