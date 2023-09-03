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

func newUserTenantRole(db *gorm.DB, opts ...gen.DOOption) userTenantRole {
	_userTenantRole := userTenantRole{}

	_userTenantRole.userTenantRoleDo.UseDB(db, opts...)
	_userTenantRole.userTenantRoleDo.UseModel(&models.UserTenantRole{})

	tableName := _userTenantRole.userTenantRoleDo.TableName()
	_userTenantRole.ALL = field.NewAsterisk(tableName)
	_userTenantRole.ID = field.NewUint64(tableName, "id")
	_userTenantRole.CreatedAt = field.NewTime(tableName, "created_at")
	_userTenantRole.UserID = field.NewUint64(tableName, "user_id")
	_userTenantRole.TenantID = field.NewUint64(tableName, "tenant_id")
	_userTenantRole.RoleID = field.NewUint64(tableName, "role_id")

	_userTenantRole.fillFieldMap()

	return _userTenantRole
}

type userTenantRole struct {
	userTenantRoleDo userTenantRoleDo

	ALL       field.Asterisk
	ID        field.Uint64 // ID
	CreatedAt field.Time   // 创建时间
	UserID    field.Uint64 // 用户ID
	TenantID  field.Uint64 // 租户ID
	RoleID    field.Uint64 // 角色ID

	fieldMap map[string]field.Expr
}

func (u userTenantRole) Table(newTableName string) *userTenantRole {
	u.userTenantRoleDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userTenantRole) As(alias string) *userTenantRole {
	u.userTenantRoleDo.DO = *(u.userTenantRoleDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userTenantRole) updateTableName(table string) *userTenantRole {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint64(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UserID = field.NewUint64(table, "user_id")
	u.TenantID = field.NewUint64(table, "tenant_id")
	u.RoleID = field.NewUint64(table, "role_id")

	u.fillFieldMap()

	return u
}

func (u *userTenantRole) WithContext(ctx context.Context) IUserTenantRoleDo {
	return u.userTenantRoleDo.WithContext(ctx)
}

func (u userTenantRole) TableName() string { return u.userTenantRoleDo.TableName() }

func (u userTenantRole) Alias() string { return u.userTenantRoleDo.Alias() }

func (u userTenantRole) Columns(cols ...field.Expr) gen.Columns {
	return u.userTenantRoleDo.Columns(cols...)
}

func (u *userTenantRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userTenantRole) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 5)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["tenant_id"] = u.TenantID
	u.fieldMap["role_id"] = u.RoleID
}

func (u userTenantRole) clone(db *gorm.DB) userTenantRole {
	u.userTenantRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userTenantRole) replaceDB(db *gorm.DB) userTenantRole {
	u.userTenantRoleDo.ReplaceDB(db)
	return u
}

type userTenantRoleDo struct{ gen.DO }

type IUserTenantRoleDo interface {
	gen.SubQuery
	Debug() IUserTenantRoleDo
	WithContext(ctx context.Context) IUserTenantRoleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserTenantRoleDo
	WriteDB() IUserTenantRoleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserTenantRoleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserTenantRoleDo
	Not(conds ...gen.Condition) IUserTenantRoleDo
	Or(conds ...gen.Condition) IUserTenantRoleDo
	Select(conds ...field.Expr) IUserTenantRoleDo
	Where(conds ...gen.Condition) IUserTenantRoleDo
	Order(conds ...field.Expr) IUserTenantRoleDo
	Distinct(cols ...field.Expr) IUserTenantRoleDo
	Omit(cols ...field.Expr) IUserTenantRoleDo
	Join(table schema.Tabler, on ...field.Expr) IUserTenantRoleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserTenantRoleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserTenantRoleDo
	Group(cols ...field.Expr) IUserTenantRoleDo
	Having(conds ...gen.Condition) IUserTenantRoleDo
	Limit(limit int) IUserTenantRoleDo
	Offset(offset int) IUserTenantRoleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTenantRoleDo
	Unscoped() IUserTenantRoleDo
	Create(values ...*models.UserTenantRole) error
	CreateInBatches(values []*models.UserTenantRole, batchSize int) error
	Save(values ...*models.UserTenantRole) error
	First() (*models.UserTenantRole, error)
	Take() (*models.UserTenantRole, error)
	Last() (*models.UserTenantRole, error)
	Find() ([]*models.UserTenantRole, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UserTenantRole, err error)
	FindInBatches(result *[]*models.UserTenantRole, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.UserTenantRole) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserTenantRoleDo
	Assign(attrs ...field.AssignExpr) IUserTenantRoleDo
	Joins(fields ...field.RelationField) IUserTenantRoleDo
	Preload(fields ...field.RelationField) IUserTenantRoleDo
	FirstOrInit() (*models.UserTenantRole, error)
	FirstOrCreate() (*models.UserTenantRole, error)
	FindByPage(offset int, limit int) (result []*models.UserTenantRole, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserTenantRoleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userTenantRoleDo) Debug() IUserTenantRoleDo {
	return u.withDO(u.DO.Debug())
}

func (u userTenantRoleDo) WithContext(ctx context.Context) IUserTenantRoleDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userTenantRoleDo) ReadDB() IUserTenantRoleDo {
	return u.Clauses(dbresolver.Read)
}

func (u userTenantRoleDo) WriteDB() IUserTenantRoleDo {
	return u.Clauses(dbresolver.Write)
}

func (u userTenantRoleDo) Session(config *gorm.Session) IUserTenantRoleDo {
	return u.withDO(u.DO.Session(config))
}

func (u userTenantRoleDo) Clauses(conds ...clause.Expression) IUserTenantRoleDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userTenantRoleDo) Returning(value interface{}, columns ...string) IUserTenantRoleDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userTenantRoleDo) Not(conds ...gen.Condition) IUserTenantRoleDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userTenantRoleDo) Or(conds ...gen.Condition) IUserTenantRoleDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userTenantRoleDo) Select(conds ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userTenantRoleDo) Where(conds ...gen.Condition) IUserTenantRoleDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userTenantRoleDo) Order(conds ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userTenantRoleDo) Distinct(cols ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userTenantRoleDo) Omit(cols ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userTenantRoleDo) Join(table schema.Tabler, on ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userTenantRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userTenantRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userTenantRoleDo) Group(cols ...field.Expr) IUserTenantRoleDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userTenantRoleDo) Having(conds ...gen.Condition) IUserTenantRoleDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userTenantRoleDo) Limit(limit int) IUserTenantRoleDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userTenantRoleDo) Offset(offset int) IUserTenantRoleDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userTenantRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserTenantRoleDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userTenantRoleDo) Unscoped() IUserTenantRoleDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userTenantRoleDo) Create(values ...*models.UserTenantRole) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userTenantRoleDo) CreateInBatches(values []*models.UserTenantRole, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userTenantRoleDo) Save(values ...*models.UserTenantRole) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userTenantRoleDo) First() (*models.UserTenantRole, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserTenantRole), nil
	}
}

func (u userTenantRoleDo) Take() (*models.UserTenantRole, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserTenantRole), nil
	}
}

func (u userTenantRoleDo) Last() (*models.UserTenantRole, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserTenantRole), nil
	}
}

func (u userTenantRoleDo) Find() ([]*models.UserTenantRole, error) {
	result, err := u.DO.Find()
	return result.([]*models.UserTenantRole), err
}

func (u userTenantRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.UserTenantRole, err error) {
	buf := make([]*models.UserTenantRole, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userTenantRoleDo) FindInBatches(result *[]*models.UserTenantRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userTenantRoleDo) Attrs(attrs ...field.AssignExpr) IUserTenantRoleDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userTenantRoleDo) Assign(attrs ...field.AssignExpr) IUserTenantRoleDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userTenantRoleDo) Joins(fields ...field.RelationField) IUserTenantRoleDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userTenantRoleDo) Preload(fields ...field.RelationField) IUserTenantRoleDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userTenantRoleDo) FirstOrInit() (*models.UserTenantRole, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserTenantRole), nil
	}
}

func (u userTenantRoleDo) FirstOrCreate() (*models.UserTenantRole, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.UserTenantRole), nil
	}
}

func (u userTenantRoleDo) FindByPage(offset int, limit int) (result []*models.UserTenantRole, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userTenantRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userTenantRoleDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userTenantRoleDo) Delete(models ...*models.UserTenantRole) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userTenantRoleDo) withDO(do gen.Dao) *userTenantRoleDo {
	u.DO = *do.(*gen.DO)
	return u
}
