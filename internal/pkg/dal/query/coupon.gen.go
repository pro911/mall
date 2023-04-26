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

	"mall/internal/pkg/dal/model"
)

func newCoupon(db *gorm.DB, opts ...gen.DOOption) coupon {
	_coupon := coupon{}

	_coupon.couponDo.UseDB(db, opts...)
	_coupon.couponDo.UseModel(&model.Coupon{})

	tableName := _coupon.couponDo.TableName()
	_coupon.ALL = field.NewAsterisk(tableName)
	_coupon.CouponID = field.NewInt32(tableName, "coupon_id")
	_coupon.Name = field.NewString(tableName, "name")
	_coupon.Satisfy = field.NewInt32(tableName, "satisfy")
	_coupon.Minus = field.NewInt32(tableName, "minus")
	_coupon.Desc = field.NewString(tableName, "desc")
	_coupon.CreatedAt = field.NewTime(tableName, "created_at")
	_coupon.UpdatedAt = field.NewTime(tableName, "updated_at")

	_coupon.fillFieldMap()

	return _coupon
}

type coupon struct {
	couponDo couponDo

	ALL       field.Asterisk
	CouponID  field.Int32
	Name      field.String
	Satisfy   field.Int32
	Minus     field.Int32
	Desc      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (c coupon) Table(newTableName string) *coupon {
	c.couponDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c coupon) As(alias string) *coupon {
	c.couponDo.DO = *(c.couponDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *coupon) updateTableName(table string) *coupon {
	c.ALL = field.NewAsterisk(table)
	c.CouponID = field.NewInt32(table, "coupon_id")
	c.Name = field.NewString(table, "name")
	c.Satisfy = field.NewInt32(table, "satisfy")
	c.Minus = field.NewInt32(table, "minus")
	c.Desc = field.NewString(table, "desc")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *coupon) WithContext(ctx context.Context) ICouponDo { return c.couponDo.WithContext(ctx) }

func (c coupon) TableName() string { return c.couponDo.TableName() }

func (c coupon) Alias() string { return c.couponDo.Alias() }

func (c *coupon) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *coupon) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 7)
	c.fieldMap["coupon_id"] = c.CouponID
	c.fieldMap["name"] = c.Name
	c.fieldMap["satisfy"] = c.Satisfy
	c.fieldMap["minus"] = c.Minus
	c.fieldMap["desc"] = c.Desc
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c coupon) clone(db *gorm.DB) coupon {
	c.couponDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c coupon) replaceDB(db *gorm.DB) coupon {
	c.couponDo.ReplaceDB(db)
	return c
}

type couponDo struct{ gen.DO }

type ICouponDo interface {
	gen.SubQuery
	Debug() ICouponDo
	WithContext(ctx context.Context) ICouponDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICouponDo
	WriteDB() ICouponDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICouponDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICouponDo
	Not(conds ...gen.Condition) ICouponDo
	Or(conds ...gen.Condition) ICouponDo
	Select(conds ...field.Expr) ICouponDo
	Where(conds ...gen.Condition) ICouponDo
	Order(conds ...field.Expr) ICouponDo
	Distinct(cols ...field.Expr) ICouponDo
	Omit(cols ...field.Expr) ICouponDo
	Join(table schema.Tabler, on ...field.Expr) ICouponDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICouponDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICouponDo
	Group(cols ...field.Expr) ICouponDo
	Having(conds ...gen.Condition) ICouponDo
	Limit(limit int) ICouponDo
	Offset(offset int) ICouponDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICouponDo
	Unscoped() ICouponDo
	Create(values ...*model.Coupon) error
	CreateInBatches(values []*model.Coupon, batchSize int) error
	Save(values ...*model.Coupon) error
	First() (*model.Coupon, error)
	Take() (*model.Coupon, error)
	Last() (*model.Coupon, error)
	Find() ([]*model.Coupon, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Coupon, err error)
	FindInBatches(result *[]*model.Coupon, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Coupon) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICouponDo
	Assign(attrs ...field.AssignExpr) ICouponDo
	Joins(fields ...field.RelationField) ICouponDo
	Preload(fields ...field.RelationField) ICouponDo
	FirstOrInit() (*model.Coupon, error)
	FirstOrCreate() (*model.Coupon, error)
	FindByPage(offset int, limit int) (result []*model.Coupon, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICouponDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c couponDo) Debug() ICouponDo {
	return c.withDO(c.DO.Debug())
}

func (c couponDo) WithContext(ctx context.Context) ICouponDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c couponDo) ReadDB() ICouponDo {
	return c.Clauses(dbresolver.Read)
}

func (c couponDo) WriteDB() ICouponDo {
	return c.Clauses(dbresolver.Write)
}

func (c couponDo) Session(config *gorm.Session) ICouponDo {
	return c.withDO(c.DO.Session(config))
}

func (c couponDo) Clauses(conds ...clause.Expression) ICouponDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c couponDo) Returning(value interface{}, columns ...string) ICouponDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c couponDo) Not(conds ...gen.Condition) ICouponDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c couponDo) Or(conds ...gen.Condition) ICouponDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c couponDo) Select(conds ...field.Expr) ICouponDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c couponDo) Where(conds ...gen.Condition) ICouponDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c couponDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICouponDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c couponDo) Order(conds ...field.Expr) ICouponDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c couponDo) Distinct(cols ...field.Expr) ICouponDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c couponDo) Omit(cols ...field.Expr) ICouponDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c couponDo) Join(table schema.Tabler, on ...field.Expr) ICouponDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c couponDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICouponDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c couponDo) RightJoin(table schema.Tabler, on ...field.Expr) ICouponDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c couponDo) Group(cols ...field.Expr) ICouponDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c couponDo) Having(conds ...gen.Condition) ICouponDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c couponDo) Limit(limit int) ICouponDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c couponDo) Offset(offset int) ICouponDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c couponDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICouponDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c couponDo) Unscoped() ICouponDo {
	return c.withDO(c.DO.Unscoped())
}

func (c couponDo) Create(values ...*model.Coupon) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c couponDo) CreateInBatches(values []*model.Coupon, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c couponDo) Save(values ...*model.Coupon) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c couponDo) First() (*model.Coupon, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coupon), nil
	}
}

func (c couponDo) Take() (*model.Coupon, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coupon), nil
	}
}

func (c couponDo) Last() (*model.Coupon, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coupon), nil
	}
}

func (c couponDo) Find() ([]*model.Coupon, error) {
	result, err := c.DO.Find()
	return result.([]*model.Coupon), err
}

func (c couponDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Coupon, err error) {
	buf := make([]*model.Coupon, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c couponDo) FindInBatches(result *[]*model.Coupon, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c couponDo) Attrs(attrs ...field.AssignExpr) ICouponDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c couponDo) Assign(attrs ...field.AssignExpr) ICouponDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c couponDo) Joins(fields ...field.RelationField) ICouponDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c couponDo) Preload(fields ...field.RelationField) ICouponDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c couponDo) FirstOrInit() (*model.Coupon, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coupon), nil
	}
}

func (c couponDo) FirstOrCreate() (*model.Coupon, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Coupon), nil
	}
}

func (c couponDo) FindByPage(offset int, limit int) (result []*model.Coupon, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c couponDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c couponDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c couponDo) Delete(models ...*model.Coupon) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *couponDo) withDO(do gen.Dao) *couponDo {
	c.DO = *do.(*gen.DO)
	return c
}
