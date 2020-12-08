package auth

import (
	"testing"
)

var (
	clusterUUID = "8b891540-650a-4f2a-839f-82e8b2cc222e"
	magnum      = "magnum"
	stoneCli    = NewKeystone("http://10.60.41.127:35357/v3", "gAAAAABfzyGfPEinFN3F29BPrMlpotvYsboB9RtZpBSbDNujUaVMArB5ccra_zIEb8fA1WY7kAYEbenJPNbMHJWx84kWH3DxocGR1Ku7oJJEkhEiYQzPNEL7jTKSRWgkVROWg8WGL0mX4kjAbhHUDW-9mJnIi4g_UKk4CQu8LYBaQopn4B0DKEI")
)

func TestCluster(t *testing.T) {
	service, err := stoneCli.Service(magnum)
	if err != nil {
		t.Fatal("can't get service", err)
	}
	endpoints, err := stoneCli.Endpoints(map[string]string{"service_id": service.ID, "interface": "public"})
	if err != nil {
		t.Fatal("can't get endpoint", err)
	}
	if len(endpoints.Endpoints) == 0 {
		t.Error("not found magnum endpoint")
		return
	}
	mag := endpoints.Endpoints[0]
	cluster, err := stoneCli.Cluster(mag.URL, clusterUUID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cluster)
}

func TestService(t *testing.T) {
	service, err := stoneCli.Service(magnum)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(service)
}

func TestEndpoint(t *testing.T) {
	service, err := stoneCli.Service(magnum)
	if err != nil {
		t.Fatal(err)
	}
	endpoints, err := stoneCli.Endpoints(map[string]string{"service_id": service.ID, "interface": "public"})
	if err != nil {
		t.Fatal(err)
	}
	for _, ep := range endpoints.Endpoints {
		t.Log(ep)
	}
}

// TestCA get root ca and create client private key、client ca
func TestCA(t *testing.T) {
	service, err := stoneCli.Service(magnum)
	if err != nil {
		t.Fatal("can't get service", err)
	}
	endpoints, err := stoneCli.Endpoints(map[string]string{"service_id": service.ID, "interface": "public"})
	if err != nil {
		t.Fatal("can't get endpoint", err)
	}
	if len(endpoints.Endpoints) == 0 {
		t.Error("not found magnum endpoint")
		return
	}
	mag := endpoints.Endpoints[0]
	cert, err := stoneCli.CA(mag.URL, clusterUUID)
	if err != nil {
		t.Fatal("can't get ca", err)
	}
	privateKey, ca, err := stoneCli.Sign(mag.URL, clusterUUID)
	if err != nil {
		t.Fatal("sign client certificate failed", err)
	}
	t.Log("root ca", cert, ",client private key", privateKey, ",client ca", ca)
}

func TestToken(t *testing.T) {
	token, err := stoneCli.Token()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}
