package query

import (
	"github.com/mojo-zd/helm-crabstick/data/types"
	"github.com/sirupsen/logrus"
)

// ReleaseDao the database operator of instance
type ReleaseDao struct {
	cornerstone
}

// NewReleaseDao ...
func NewReleaseDao() *ReleaseDao {
	return &ReleaseDao{}
}

// Create create instance
func (i *ReleaseDao) Create(o *types.Release) error {
	return i.cornerstone.Create(o)
}

// CreateBathes you can create multi object if param slice is a slice or array
// otherwise create a single object
func (i *ReleaseDao) CreateBathes(slice interface{}) error {
	return i.cornerstone.CreateBathes(slice)
}

// Get find one object, you can spec condition with o's attr
func (i *ReleaseDao) Get(o *types.Release) error {
	return i.cornerstone.Get(o)
}

// List find all record witch match the condition
func (i *ReleaseDao) List(condition *types.Release) ([]*types.Release, error) {
	out := []*types.Release{}
	err := i.cornerstone.List(&out, &types.Release{}, condition)
	return out, err
}

// Update update object if you want to update bool to false you must call UpdateWithCols method
func (i *ReleaseDao) Update(o interface{}) error {
	return i.cornerstone.Update(o)
}

// UpdateWithCols update object with spec attr
func (i *ReleaseDao) UpdateWithCols(condition, values map[string]interface{}) error {
	if values == nil {
		logrus.Warnln("skip update action because values is nil")
		return nil
	}
	return i.cornerstone.UpdateWithCols(&types.Release{}, condition, values)
}

// Delete delete types.Release with id
func (i *ReleaseDao) Delete(release *types.Release) error {
	// avoid to delete all data
	if release.ID == 0 {
		logrus.Warnf("skip to delete release, because release's id is 0")
		return nil
	}
	return i.cornerstone.Delete(release)
}
