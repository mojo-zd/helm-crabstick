package page

import "reflect"

type Pagination struct {
	Records  interface{} `json:"records"`
	Total    int64       `json:"total"`
	Pages    int64       `json:"pages"`
	PageSize int64       `json:"pageSize"`
	Current  int64       `json:"current"`
}

func NewPage(size, index int64, data interface{}) *Pagination {
	p := &Pagination{
		PageSize: size,
		Current:  index,
	}
	typ := reflect.TypeOf(data)
	if typ.Kind() != reflect.Slice {
		p.Records = data
	}
	// ---
	data
	return p
}

func (pagination *Pagination) Offset() int64 {
	if pagination.Current == 0 {
		return 0
	}

	return (pagination.Current - 1) * pagination.PageSize
}

func (pagination *Pagination) SetTotal(total int64) {
	if total == 0 {
		return
	}

	if pagination.PageSize == 0 {
		panic("page size must positive number")
	}

	pagination.Total = total
	if pagination.Total%pagination.PageSize > 0 {
		pagination.Pages = pagination.Total/pagination.PageSize + 1
	} else {
		pagination.Pages = pagination.Total / pagination.PageSize
	}
}
