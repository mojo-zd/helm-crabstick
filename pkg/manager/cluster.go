package manager

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/mojo-zd/helm-crabstick/pkg/util/file"

	"github.com/mojo-zd/helm-crabstick/pkg/auth"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	magnum   = "magnum"
	inter    = "public"
	caFile   = "ca.crt"
	certFile = "client.crt"
	keyFile  = "client.key"
)

var authPath = fmt.Sprintf("%s/.certs/clusters", file.HomeDir())

type Manager interface {
	// instance kubernetes clientset
	Client(clusterUUID, token string) (Cluster, error)
}

type Cluster struct {
	UUID       string
	ApiAddress string // e.g. ip:port
	CAFile     string
	CertFile   string
	KeyFile    string
	CAData     string
	CertData   string
	KeyData    string
	Client     kubernetes.Interface
}

type clusterManager struct {
	keystone string // keystone address
	cache    map[string]*Cluster
	lock     sync.RWMutex
}

func NewClusterManager(keystoneAddr string) Manager {
	return &clusterManager{
		keystone: keystoneAddr,
		cache:    make(map[string]*Cluster),
	}
}

// Client create client witch interacts with kubernetes
func (mgr *clusterManager) Client(clusterUUID, token string) (Cluster, error) {
	var out Cluster
	stoneCli := auth.NewKeystone(mgr.keystone, token)
	service, err := stoneCli.Service(magnum)
	if err != nil {
		logrus.Errorf("can't get service[%s], err:%s", magnum, err)
		return out, err
	}
	endpoints, err := stoneCli.Endpoints(map[string]string{"service_id": service.ID, "interface": inter})
	if err != nil {
		logrus.Errorf("can't get endpoint[%s], filter is service_id[%s]、 interface[%s], err:%s", magnum, clusterUUID, inter, err)
		return out, err
	}
	if len(endpoints.Endpoints) == 0 {
		logrus.Error("not found magnum endpoint")
		return out, err
	}
	mag := endpoints.Endpoints[0]
	cluster, err := stoneCli.Cluster(mag.URL, clusterUUID)
	if err != nil {
		logrus.Errorf("can't get cluster[%s], err:%s", clusterUUID, err.Error())
		return out, err
	}
	if cluster.HealthStatus != auth.HealthStatus {
		logrus.Warnf("cluster status unhealthy")
		return out, errors.New("cluster status unhealthy")
	}

	// check object from cache and return client
	if val, ok := mgr.cache[cluster.UUID]; ok {
		logrus.Debugf("get cluster's[%s] kubernetes client from cache", clusterUUID)
		return *val, nil
	}

	ca, err := stoneCli.CA(mag.URL, clusterUUID)
	if err != nil {
		logrus.Errorf("can't get cluster[%s]'s ca, err:%s", clusterUUID, err.Error())
		return out, err
	}

	privateKey, cert, err := stoneCli.Sign(mag.URL, clusterUUID)
	if err != nil {
		logrus.Errorf("sign client certificate failed of cluster[%s], err:%s", clusterUUID, err.Error())
		return out, err
	}
	out.UUID = cluster.UUID
	out.CAData = ca.PEM
	out.CertData = cert
	out.KeyData = privateKey
	out.ApiAddress = cluster.ApiAddress
	if err = out.generateAuthFile(); err != nil {
		return out, err
	}

	client, err := kubernetes.NewForConfig(&rest.Config{
		Host: cluster.ApiAddress,
		TLSClientConfig: rest.TLSClientConfig{
			CAFile:   out.CAFile,
			CertFile: out.CertFile,
			KeyFile:  out.KeyFile,
		},
	})
	if err != nil {
		logrus.Errorf("new kubernetes client instance failed of cluster[%s], err:%s", clusterUUID, err)
		return out, err
	}
	out.Client = client
	mgr.cache[cluster.UUID] = &out
	return out, err
}

// generateAuthFile generate auth file to spec dir
func (c *Cluster) generateAuthFile() error {
	rootPath := path.Join(authPath, c.UUID)
	if err := os.MkdirAll(rootPath, 0775); err != nil {
		logrus.Errorf("make dir[%s] failed, err: %s", rootPath, err.Error())
		return err
	}

	caPath := path.Join(rootPath, caFile)
	certPath := path.Join(rootPath, certFile)
	keyPath := path.Join(rootPath, keyFile)
	if err := ioutil.WriteFile(caPath, []byte(c.CAData), 0755); err != nil {
		logrus.Errorf("write %s failed, err: %s", caPath, err.Error())
		return err
	}

	if err := ioutil.WriteFile(certPath, []byte(c.CertData), 0755); err != nil {
		logrus.Errorf("write %s failed, err: %s", certPath, err.Error())
		return err
	}

	if err := ioutil.WriteFile(keyPath, []byte(c.KeyData), 0755); err != nil {
		logrus.Errorf("write %s failed, err: %s", keyPath, err.Error())
		return err
	}
	c.CAFile = caPath
	c.CertFile = certPath
	c.KeyFile = keyPath
	return nil
}
