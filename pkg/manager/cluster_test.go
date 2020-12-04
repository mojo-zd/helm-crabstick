package manager

import (
	"context"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	clusterUUID = "8b891540-650a-4f2a-839f-82e8b2cc222e"
	keystone    = "http://10.60.41.127:35357/v3"
	token       = "gAAAAABfyEKUWsORyjwKYTa5ZT_jH0_xITTuf4d0R7jPmymO0JTB7JAWVxHwTU7Ys4TWeTLXGdn5qg0iJcIzPAKxI8n96UddWslP9hoizF8jFJ4bVB2hIE5JuTlqC-YxVRrwg6V25280BgP71L-mCxKDICXzz_kQwA"
)

func TestClusterManager(t *testing.T) {
	manager := NewClusterManager(keystone)
	cluster, err := manager.Client(clusterUUID, token)
	if err != nil {
		t.Fatal(err)
	}
	namespaces, err := cluster.Client.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		t.Fatal("can't list namespaces", err)
	}
	t.Log(namespaces.Items)
}
