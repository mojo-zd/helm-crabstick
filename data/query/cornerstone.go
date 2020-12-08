package query

import (
	"reflect"

	"github.com/mojo-zd/helm-crabstick/data/conn"
)

// cornerstone the abstract of database operation
type cornerstone struct {
}

// Create create instance to db
func (c *cornerstone) Create(o interface{}) error {
	return conn.GetDB().Create(o).Error
}

// CreateBathes create object with bathes
func (c *cornerstone) CreateBathes(slice interface{}) error {
	len := 1 // the bath size
	val := reflect.ValueOf(slice)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		len = val.Len()
	}
	return conn.GetDB().CreateInBatches(slice, len).Error
}

// Get get instance from db
func (c *cornerstone) Get(out interface{}) error {
	return conn.GetDB().Where(out).Take(out).Error
}

// List find all record witch matches
// only support struct type
func (c *cornerstone) List(out, table, condition interface{}) error {
	db := conn.GetDB().Model(table)
	if condition != nil {
		db = db.Where(condition)
	}
	return db.Find(out).Error
}

// Update update object if you want to update bool to false you must call UpdateWithCols method
func (c *cornerstone) Update(o interface{}) error {
	return conn.GetDB().Updates(o).Error
}

// UpdateWithCols spec the cols to update object
func (c *cornerstone) UpdateWithCols(model interface{}, condition, cols map[string]interface{}) error {
	return conn.GetDB().Model(model).Where(condition).UpdateColumns(cols).Error
}

// DeleteBathes delete obj with condition, will delete all if condition is nil
func (c *cornerstone) DeleteBathes(model interface{}, condition map[string]interface{}) error {
	return conn.GetDB().Where(condition).Delete(model).Error
}

// Delete delete with id
func (c *cornerstone) Delete(o interface{}) error {
	return conn.GetDB().Delete(o).Error
}
