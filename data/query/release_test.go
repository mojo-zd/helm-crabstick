package query

import (
	"testing"
	"time"

	"github.com/mojo-zd/helm-crabstick/data/conn"
	"github.com/mojo-zd/helm-crabstick/data/types"
)

func init() {
	conn.NewConn(conn.Config{
		Host:        "127.0.0.1",
		Port:        "3307",
		Username:    "root",
		Password:    "root123",
		Schema:      "app_market",
		MaxConn:     4,
		MaxIdleConn: 2,
		MaxIdleTime: 3600,
	})
}

var (
	releaseDao = NewReleaseDao()
)

func TestCreateRelease(t *testing.T) {
	inst := &types.Release{Project: "tenant1", Domain: "default"}
	if err := releaseDao.Create(inst); err != nil {
		t.Fatal(err)
	}
	t.Log(inst)
}

func TestBathesCreateRelease(t *testing.T) {
	cases := []struct {
		object interface{}
		want   error
	}{
		{
			object: &[]*types.Release{
				{
					Project: "t1", Domain: "default",
				},
				{
					Project: "t2", Domain: "default",
				},
			},
			want: nil,
		},
		{
			object: &types.Release{
				Project: "single",
				Domain:  "default",
			},
			want: nil,
		},
	}

	for _, cas := range cases {
		if err := releaseDao.CreateBathes(cas.object); err != cas.want {
			t.Fatal(err)
		}
		t.Log(cas.object)
	}
}

func TestFindRelease(t *testing.T) {
	out := &types.Release{CreatedAt: 1607322668}
	if err := releaseDao.Get(out); err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

func TestFindAllRelease(t *testing.T) {
	out, err := releaseDao.List(&types.Release{CreatedAt: 1607322668})
	if err != nil {
		t.Fatal(err)
	}
	for _, o := range out {
		t.Log(o)
	}
}

func TestUpdateRelease(t *testing.T) {
	t.Log(time.Now().Unix())
	if err := releaseDao.Update(&types.Release{ID: 2, Project: "default"}); err != nil {
		t.Fatal(err)
	}
	if err := releaseDao.UpdateWithCols(map[string]interface{}{"id": 1},
		map[string]interface{}{"public": true}); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteRelease(t *testing.T) {
	if err := releaseDao.Delete(&types.Release{ID: 3}); err != nil {
		t.Fatal(err)
	}
	t.Log("delete success")
}
