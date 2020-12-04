package page

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

var (
	DefSize    int64 = 10
	DefCurrent int64 = 1
)

type Pagination struct {
	Rows     interface{} `json:"rows"`
	Total    int64       `json:"total"`
	PageSize int64       `json:"pageSize"`
	Current  int64       `json:"current"`
}

func NewPagination(data interface{}, size, current int64) *Pagination {
	if size <= 0 {
		size = DefSize
	}

	if current <= 0 {
		current = DefCurrent
	}

	page := &Pagination{PageSize: size, Current: current}
	if data == nil {
		return page
	}
	page.section(data)
	return page
}

func (page *Pagination) section(data interface{}) {
	ty := reflect.TypeOf(data)
	value := reflect.ValueOf(data)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
		value = value.Elem()
	}

	if ty.Kind() != reflect.Slice {
		logrus.Warnln("data is not a slice")
		page.Rows = data
		page.Total = DefCurrent
		return
	}

	var pages int64
	len := int64(value.Len())
	if len%page.PageSize == 0 {
		pages = len / page.PageSize
	} else {
		pages = len/page.PageSize + 1
	}

	// setting the last page if current great than max page
	if pages < page.Current {
		page.Current = pages
	}

	start := (page.Current - 1) * page.PageSize
	end := page.Current * page.PageSize
	// return the last
	if len < end {
		page.Rows = value.Slice(int(start), int(len)).Interface()
	} else {
		page.Rows = value.Slice(int(start), int(end)).Interface()
	}
	return
}
