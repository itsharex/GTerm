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

	"github.com/MisakaTAT/GTerm/backend/dal/model"
)

func newMetadata(db *gorm.DB, opts ...gen.DOOption) metadata {
	_metadata := metadata{}

	_metadata.metadataDo.UseDB(db, opts...)
	_metadata.metadataDo.UseModel(&model.Metadata{})

	tableName := _metadata.metadataDo.TableName()
	_metadata.ALL = field.NewAsterisk(tableName)
	_metadata.ID = field.NewUint(tableName, "id")
	_metadata.CreatedAt = field.NewTime(tableName, "created_at")
	_metadata.UpdatedAt = field.NewTime(tableName, "updated_at")
	_metadata.DeletedAt = field.NewField(tableName, "deleted_at")
	_metadata.HostID = field.NewUint(tableName, "host_id")
	_metadata.Processors = field.NewUint(tableName, "processors")
	_metadata.MemTotal = field.NewUint(tableName, "mem_total")
	_metadata.OS = field.NewString(tableName, "os")

	_metadata.fillFieldMap()

	return _metadata
}

type metadata struct {
	metadataDo

	ALL        field.Asterisk
	ID         field.Uint
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	HostID     field.Uint
	Processors field.Uint
	MemTotal   field.Uint
	OS         field.String

	fieldMap map[string]field.Expr
}

func (m metadata) Table(newTableName string) *metadata {
	m.metadataDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m metadata) As(alias string) *metadata {
	m.metadataDo.DO = *(m.metadataDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *metadata) updateTableName(table string) *metadata {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewUint(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.HostID = field.NewUint(table, "host_id")
	m.Processors = field.NewUint(table, "processors")
	m.MemTotal = field.NewUint(table, "mem_total")
	m.OS = field.NewString(table, "os")

	m.fillFieldMap()

	return m
}

func (m *metadata) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *metadata) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 8)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["host_id"] = m.HostID
	m.fieldMap["processors"] = m.Processors
	m.fieldMap["mem_total"] = m.MemTotal
	m.fieldMap["os"] = m.OS
}

func (m metadata) clone(db *gorm.DB) metadata {
	m.metadataDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m metadata) replaceDB(db *gorm.DB) metadata {
	m.metadataDo.ReplaceDB(db)
	return m
}

type metadataDo struct{ gen.DO }

type IMetadataDo interface {
	gen.SubQuery
	Debug() IMetadataDo
	WithContext(ctx context.Context) IMetadataDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMetadataDo
	WriteDB() IMetadataDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMetadataDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMetadataDo
	Not(conds ...gen.Condition) IMetadataDo
	Or(conds ...gen.Condition) IMetadataDo
	Select(conds ...field.Expr) IMetadataDo
	Where(conds ...gen.Condition) IMetadataDo
	Order(conds ...field.Expr) IMetadataDo
	Distinct(cols ...field.Expr) IMetadataDo
	Omit(cols ...field.Expr) IMetadataDo
	Join(table schema.Tabler, on ...field.Expr) IMetadataDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMetadataDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMetadataDo
	Group(cols ...field.Expr) IMetadataDo
	Having(conds ...gen.Condition) IMetadataDo
	Limit(limit int) IMetadataDo
	Offset(offset int) IMetadataDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMetadataDo
	Unscoped() IMetadataDo
	Create(values ...*model.Metadata) error
	CreateInBatches(values []*model.Metadata, batchSize int) error
	Save(values ...*model.Metadata) error
	First() (*model.Metadata, error)
	Take() (*model.Metadata, error)
	Last() (*model.Metadata, error)
	Find() ([]*model.Metadata, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Metadata, err error)
	FindInBatches(result *[]*model.Metadata, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Metadata) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMetadataDo
	Assign(attrs ...field.AssignExpr) IMetadataDo
	Joins(fields ...field.RelationField) IMetadataDo
	Preload(fields ...field.RelationField) IMetadataDo
	FirstOrInit() (*model.Metadata, error)
	FirstOrCreate() (*model.Metadata, error)
	FindByPage(offset int, limit int) (result []*model.Metadata, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMetadataDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m metadataDo) Debug() IMetadataDo {
	return m.withDO(m.DO.Debug())
}

func (m metadataDo) WithContext(ctx context.Context) IMetadataDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m metadataDo) ReadDB() IMetadataDo {
	return m.Clauses(dbresolver.Read)
}

func (m metadataDo) WriteDB() IMetadataDo {
	return m.Clauses(dbresolver.Write)
}

func (m metadataDo) Session(config *gorm.Session) IMetadataDo {
	return m.withDO(m.DO.Session(config))
}

func (m metadataDo) Clauses(conds ...clause.Expression) IMetadataDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m metadataDo) Returning(value interface{}, columns ...string) IMetadataDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m metadataDo) Not(conds ...gen.Condition) IMetadataDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m metadataDo) Or(conds ...gen.Condition) IMetadataDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m metadataDo) Select(conds ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m metadataDo) Where(conds ...gen.Condition) IMetadataDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m metadataDo) Order(conds ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m metadataDo) Distinct(cols ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m metadataDo) Omit(cols ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m metadataDo) Join(table schema.Tabler, on ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m metadataDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m metadataDo) RightJoin(table schema.Tabler, on ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m metadataDo) Group(cols ...field.Expr) IMetadataDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m metadataDo) Having(conds ...gen.Condition) IMetadataDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m metadataDo) Limit(limit int) IMetadataDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m metadataDo) Offset(offset int) IMetadataDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m metadataDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMetadataDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m metadataDo) Unscoped() IMetadataDo {
	return m.withDO(m.DO.Unscoped())
}

func (m metadataDo) Create(values ...*model.Metadata) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m metadataDo) CreateInBatches(values []*model.Metadata, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m metadataDo) Save(values ...*model.Metadata) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m metadataDo) First() (*model.Metadata, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metadata), nil
	}
}

func (m metadataDo) Take() (*model.Metadata, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metadata), nil
	}
}

func (m metadataDo) Last() (*model.Metadata, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metadata), nil
	}
}

func (m metadataDo) Find() ([]*model.Metadata, error) {
	result, err := m.DO.Find()
	return result.([]*model.Metadata), err
}

func (m metadataDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Metadata, err error) {
	buf := make([]*model.Metadata, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m metadataDo) FindInBatches(result *[]*model.Metadata, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m metadataDo) Attrs(attrs ...field.AssignExpr) IMetadataDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m metadataDo) Assign(attrs ...field.AssignExpr) IMetadataDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m metadataDo) Joins(fields ...field.RelationField) IMetadataDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m metadataDo) Preload(fields ...field.RelationField) IMetadataDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m metadataDo) FirstOrInit() (*model.Metadata, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metadata), nil
	}
}

func (m metadataDo) FirstOrCreate() (*model.Metadata, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Metadata), nil
	}
}

func (m metadataDo) FindByPage(offset int, limit int) (result []*model.Metadata, count int64, err error) {
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

func (m metadataDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m metadataDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m metadataDo) Delete(models ...*model.Metadata) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *metadataDo) withDO(do gen.Dao) *metadataDo {
	m.DO = *do.(*gen.DO)
	return m
}
