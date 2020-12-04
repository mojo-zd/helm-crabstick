package page

import "testing"

var Data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

func TestPage(t *testing.T) {
	var pageSize int64 = 10
	cases := []struct {
		current int64
		want    int64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
	}

	for _, cas := range cases {
		page := NewPagination(Data, pageSize, cas.current)
		if page.Current == cas.want {
			t.Log("current:", cas.current, page.Rows)
			continue
		}
		t.Logf("current is %d, expect is %d", cas.current, cas.want)
	}
}
