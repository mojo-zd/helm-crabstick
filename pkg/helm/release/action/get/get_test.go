package get

import (
	"fmt"
	"os"
	"path"
	"testing"

	"k8s.io/client-go/rest"

	"github.com/mojo-zd/helm-crabstick/pkg/manager"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/manager/kube"
	"github.com/mojo-zd/helm-crabstick/pkg/helm/util"
	"k8s.io/client-go/kubernetes"
)

var (
	home = os.Getenv("HOME")
	conf = config.Config{
		Repository: &config.Repository{
			Name:  "bitnami",
			URL:   "https://charts.bitnami.com/bitnami",
			Cache: fmt.Sprintf("%s/.cache/helm", home),
		},
		KubeConf: fmt.Sprintf("%s/.kube/config", home),
	}
	namespace   = "kubeapps"
	releaseName = "kubeapps"
)

var (
	ca = `-----BEGIN CERTIFICATE-----
MIID1DCCArygAwIBAgIUO1701P0pXWQ19YtEqjN9F8eYACAwDQYJKoZIhvcNAQEL
BQAwbzELMAkGA1UEBhMCQ04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UEBxMI
U2hlbnpoZW4xEzARBgNVBAoTCkt1YmVybmV0ZXMxDzANBgNVBAsTBldpc2UyQzET
MBEGA1UEAxMKa3ViZXJuZXRlczAgFw0yMDExMjAwMzQyMDBaGA8yMTIwMTAyNzAz
NDIwMFowbzELMAkGA1UEBhMCQ04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UE
BxMIU2hlbnpoZW4xEzARBgNVBAoTCkt1YmVybmV0ZXMxDzANBgNVBAsTBldpc2Uy
QzETMBEGA1UEAxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC
AQoCggEBAKk2rHiHPLEZhg/crAWbk0DyT1eUOb3Pnk8EtTSEi70nrKtp26/0qFq0
+XZaAKXgyM1HeFa4lzD1avt2Xq8mT5eG9N2lqV8XiOGN24QI9KDsda+ZwzIbgzzT
z+z35D99t4ixm9hR0r2tHixXGFHcGlqBLWmRegKPdafQCBDpBwB8h3iTbnm6ciQZ
Nv/x3Jyr9y6EJpJvTBKq+Ip2TY2vVDPXHwSKrfdb5wbvmF+EZM843vPWkw22YZGP
bCG5WwsChNhspsJ10eyxoqlUbf13+SQEDHGQd/NOSANU8ajpjQEG0tuFUw9nyDR+
bi6RWMPuT6/gZoO+xSxhurRFLU3IChECAwEAAaNmMGQwDgYDVR0PAQH/BAQDAgEG
MBIGA1UdEwEB/wQIMAYBAf8CAQIwHQYDVR0OBBYEFJheUYhyzyoD9UCx+prm+A2O
E5jwMB8GA1UdIwQYMBaAFJheUYhyzyoD9UCx+prm+A2OE5jwMA0GCSqGSIb3DQEB
CwUAA4IBAQA+CSoUsytcM1PuVYfGewSVXFQqKe2fLTowY0myF5Wqbnm9vytn+G1f
eH+gkIJHrqVQfBIV5qkhN9YLK16C8tpF0GWME2kCpI/5Bin+0lysvlvj2FAn6rV4
syfG3RY2+LJif3Si+5QVbvz2jIz/MLP6csfY7KnCIeWHJyW8DKCF0R6W3y24vQDC
RKhyMRVlRNbCygwiMfxYk8Wqr/mRYbLBLxV13PkGDNnOMA2xzmsZAktQ1d5LzTs3
cJSfeEPiBDvVbXLoarohTZ+sWQFsJvqLfAXp+dr9dIdAqCALqGdu6S5+/XhVip3b
oP6TlFqAHahcvS+T3Jthg69i6spXJFaO
-----END CERTIFICATE-----`
	cert = `-----BEGIN CERTIFICATE-----
MIID9zCCAt+gAwIBAgIUB5BCnGsnCSgOZQ7XQ9v5Yzzbsh8wDQYJKoZIhvcNAQEL
BQAwbzELMAkGA1UEBhMCQ04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UEBxMI
U2hlbnpoZW4xEzARBgNVBAoTCkt1YmVybmV0ZXMxDzANBgNVBAsTBldpc2UyQzET
MBEGA1UEAxMKa3ViZXJuZXRlczAgFw0yMDExMjAwMzQyMDBaGA8yMDcwMTEwODAz
NDIwMFoweTELMAkGA1UEBhMCQ04xEjAQBgNVBAgTCUd1YW5nZG9uZzERMA8GA1UE
BxMIU2hlbnpoZW4xFzAVBgNVBAoTDnN5c3RlbTptYXN0ZXJzMQ8wDQYDVQQLEwZX
aXNlMkMxGTAXBgNVBAMTEGt1YmVybmV0ZXMtYWRtaW4wggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQDOOjI6bQWcbwHO7iucBviAcMeRYL46xhRVqESdHDwv
srLZZ/KdB87HUXY7V+gS6X1D0Gh0DhbxnpdTIJ5tCNPs6HytE1h3wGemxeJLOMCD
+9PnqKpXC+1jf/LS7J4wHH2Yw9NXGZne1ViztVzkChmnFOXddDUB24f77fOUmplO
cFggXSYUAJXqbDKs0JWQ+111Dlqz5R1Dq7+4JScv5MQO9PYo2C949gHj7NhJIcGN
zLInjYntrxfMnjXtTzJ/bytrdZvo32lFWqm3uOhlubajlukg466gaN9loe2ZR57C
+NQose4uZq7o6EHDUNNM57LfeTQswoN0shxyuI4MKJF7AgMBAAGjfzB9MA4GA1Ud
DwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIwDAYDVR0T
AQH/BAIwADAdBgNVHQ4EFgQUVX/UhSz6Sm2Az4Gkd5HL/iRa2NEwHwYDVR0jBBgw
FoAUmF5RiHLPKgP1QLH6mub4DY4TmPAwDQYJKoZIhvcNAQELBQADggEBAE3a4GT9
NKgh+ol4MhQzqT2usimZisIeIFexoc7SvrFTKZfj42Wc1uoalQh43XrhLeaES2pS
uATOwLiMSnhrdHTj0p0ytDEKAy0Surfg95wwIN5oEfzGu8GFn30T8TrwskrdwkAw
F20NMJzmMPjbsZx3XRu/SLtlOjzo4bQZTczZEXuZDne33MvJBqxzlETVJaYxrqgc
SnOGuijJgcDLH7f4flO8UhtIFiAbwPmpixBdml+bjJ3cnIzvvXeoBMbWD2YhRPUk
m23IPzsgmilFliEJJaghHwd0W74n45XSCI4jw595NShgLOqZHzxNZZIy1/vAqEMu
g05UVsZxzmen4J4=
-----END CERTIFICATE-----`
	key = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAzjoyOm0FnG8Bzu4rnAb4gHDHkWC+OsYUVahEnRw8L7Ky2Wfy
nQfOx1F2O1foEul9Q9BodA4W8Z6XUyCebQjT7Oh8rRNYd8BnpsXiSzjAg/vT56iq
VwvtY3/y0uyeMBx9mMPTVxmZ3tVYs7Vc5AoZpxTl3XQ1AduH++3zlJqZTnBYIF0m
FACV6mwyrNCVkPtddQ5as+UdQ6u/uCUnL+TEDvT2KNgvePYB4+zYSSHBjcyyJ42J
7a8XzJ417U8yf28ra3Wb6N9pRVqpt7joZbm2o5bpIOOuoGjfZaHtmUeewvjUKLHu
Lmau6OhBw1DTTOey33k0LMKDdLIccriODCiRewIDAQABAoIBADpYwMlC+yltRsez
HueAGWLNhckd4/RhAnPRrcf9qxGbr3pPLJc9FEXUSG01y9U99lDvb/4V1mv6ALpm
KiyTKNKIXG3jYU5QQ4MtzX6WyfENmMCcOcVy/HEATEVc6MyX4vkLvomQFrazCeue
Tm++Y8+f3AEx1aV25RxkEFxk+Sb1o9g5kqrHP4tmvdrAtDxwxXiKXsjX8ionGihN
Ga4zccY6JFpYgmXTAddo88kStmUigx3yWsGPSZYeuRyecpdpjjY8We+e7LEX+Ti0
WtQDEtVqxJL5BmVZKjCvhWyiaH1Hj2P2+SvT8OzkWRsZfWPx7Nf6Q8fe9kykccvi
2x4TeCECgYEA2EwJFcsGqSgjPjEECHjXUoKQc3mSwe/I/yvktL3/X1RdAcwsOOqU
Me06UbQP+rhumEqU/VAIqIkKWz/1yKyYl8HbbwDluOzXohJX8eejkqQhWL56Xd/M
CFdNWnpxj0V4K7tfG7p4TK/f/DX2CYhXAM9kpeS6OZ+0zGLyqFOIvHkCgYEA9BT6
eYp5tpnT+BSL3i2G5gKz7G4O8ITvwOpBsJyD1k49ykXoHx10A5ELAkKxaiqYqimM
i5FVilU3zZkBLFTWjewA8tYgHy8cXEU8dwaiUfBqaJHwWgZzs/1TIyXMBX+n2+6M
GnTay7FhNpDhgVG0Y2UOybryhurZRZBpthqBGJMCgYAy2i2IoiL+wiEHDh8UntSA
4ZF0lLCcR/PJilhK5iCUGRGEyqva9cvBsTR04RCgsZvO0joVFCv088MrkO4IMAvw
IfOlNWDNCWHpCMcEaKFcaJoucxnx2BvwGhZln0PzmzGVlofVzRFbdj4C3ezqcNOD
rT7MgeoGgjXPl7PVP052gQKBgQCNZP24nORnSHOHufdQjNUht50dMKCM6qWs/sdx
FSo2YnrfC2ItbDWBv2s+Mv5tvyFTKeCWFWoVScqa2rDYSolEC9x80FgpWHQ4a49c
cEZl6zzpOOmgbS5nrS+VI9cttEa8XFNjHCCHcUkcgA9yh69VCPzpFdhbGf8lkkP6
zx3L6wKBgQDW//6LEBcfVBNlkFJ/ipXuwBQ6BVh6KoMdd+ORz7ftkm6HsjYCKk5M
CCpt3TkcTO13AFQ/YQqQBqcOFjltp3+wnm8k9UVD8krB+t8whZ41OeongWkl8Ggf
fVrKWymCejRtNYhxd8CroajUmssRqnFaeF6lqIBIrncW3wZTChSdHg==
-----END RSA PRIVATE KEY-----`
	host = "https://master-01:6443"
)

func TestReleaseList(t *testing.T) {
	root := path.Join(home, ".certs/clusters/8b891540-650a-4f2a-839f-82e8b2cc222e/")
	cf := &rest.Config{
		Host: host,
		TLSClientConfig: rest.TLSClientConfig{
			CAFile:   path.Join(root, "ca.crt"),
			CertFile: path.Join(root, "client.crt"),
			KeyFile:  path.Join(root, "client.key"),
		},
	}
	cluster := &manager.Cluster{
		UUID:       "8b891540-650a-4f2a-839f-82e8b2cc222e",
		ApiAddress: host,
		CAFile:     cf.CAFile,
		CertFile:   cf.CertFile,
		KeyFile:    cf.KeyFile,
	}
	cli, err := kubernetes.NewForConfig(cf)
	if err != nil {
		t.Fatal("new reset config", err)
	}
	cluster.Client = cli
	rls, err := NewGetter(conf, cluster, kube.NewApiManager(cluster.Client)).List("", util.ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rls[0])
	//restConf, err := conf.ConfigFlags().ToRESTConfig()
	//if err != nil {
	//	t.Fatal(err)
	//	return
	//}
	//client, err := kubernetes.NewForConfig(restConf)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//getter := NewGetter(conf, client, nil)
	//releases, err := getter.List(namespace, util.ListOptions{Annotation: map[string]string{"author": "mojo"}})
	//if err != nil {
	//	t.Fatal(err)
	//	return
	//}
	//for _, release := range releases {
	//	t.Log(release.Name, release.Chart.Metadata.Annotations)
	//}
}

func TestReleaseGet(t *testing.T) {
	//client, err := getConfAndClient()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//release, err := NewGetter(conf, client, nil).Get(releaseName, namespace)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//t.Log(release.Manifest)
}

func TestReleaseKind(t *testing.T) {
	//client, err := getConfAndClient()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//getter := NewGetter(conf, client, nil)
	//rels, err := getter.List(namespace, util.ListOptions{})
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//for _, rel := range rels {
	//	t.Logf("release [%s], kind: %s", rel.Name, getter.Kind(rel.Name, namespace))
	//}
}

func TestHistory(t *testing.T) {
	//client, err := getConfAndClient()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//getter := NewGetter(conf, client, nil)
	//out, err := getter.History(releaseName, namespace)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//
	//for _, re := range out {
	//	t.Log(re.Chart, re.Description, re.AppVersion, re.Status, re.Revision, re.Updated)
	//}
}

func getConfAndClient() (kubernetes.Interface, error) {
	//restConf, _ := conf.ConfigFlags().ToRESTConfig()
	//client, err := kubernetes.NewForConfig(restConf)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}
