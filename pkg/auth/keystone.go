package auth

import (
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	pt "path"

	resty "github.com/go-resty/resty/v2"
	"github.com/mojo-zd/helm-crabstick/pkg/util/encrypt"
	"github.com/sirupsen/logrus"
)

// API
const (
	services     = "/services"
	endpoints    = "/endpoints"
	certificates = "/certificates"
)

type keystone struct {
	token   string
	address string
}

func NewKeystone(address, token string) *keystone {
	return &keystone{token: token, address: address}
}

// Sign create client certificate and return client private key, client ca
func (k *keystone) Sign(magnumURL, cluster string) (string, string, error) {
	var privateKey, cert string
	req := resty.New().R()
	// set header info for request
	k.setAuth(req).withJsonContentType(req)
	u := k.combine([]string{magnumURL, certificates}, nil)
	key, err := encrypt.GenKey()
	if err != nil {
		logrus.Errorln("can't create rsa private key", err)
		return "", "", err
	}

	csr, err := encrypt.GenCSR(key)
	if err != nil {
		logrus.Errorln("can't create csr")
		return "", "", nil
	}
	var csrStr string
	if csrStr, err = encrypt.X509ToString(encrypt.CertRequestType, csr); err != nil {
		logrus.Errorln("csr to string failed", err.Error())
		return "", "", err
	}

	privateKey, err = encrypt.X509ToString(encrypt.RsaPrivateKeyType, x509.MarshalPKCS1PrivateKey(key))
	if err != nil {
		logrus.Errorln("private key to string failed", err.Error())
		return "", "", err
	}

	resp, err := req.SetBody(map[string]interface{}{
		"bay_uuid": cluster,
		"csr":      csrStr,
	}).Post(u.String())
	if err != nil {
		logrus.Errorf("request[%s] occur exception, err:%s", u.String(), err.Error())
		return privateKey, cert, err
	}

	cer := Certificate{}
	if err = json.Unmarshal(resp.Body(), &cer); err != nil {
		logrus.Errorf("can't unmarshal to Certificate, body:%s, err:%s", resp.Body(), err.Error())
		return "", "", err
	}
	return privateKey, cer.PEM, nil
}

// CA get ca from magnum
func (k *keystone) CA(magnumURL, cluster string) (Certificate, error) {
	out := Certificate{}
	req := resty.New().R()
	// set header info for request
	k.setAuth(req).withJsonContentType(req)
	u := k.combine([]string{magnumURL, certificates, cluster}, nil)
	logrus.Debugln("request url is", u.String())
	resp, err := req.Get(u.String())
	if err != nil {
		logrus.Errorf("request[%s] occur exception, err:%s", u.String(), err.Error())
		return out, err
	}

	if err = json.Unmarshal(resp.Body(), &out); err != nil {
		logrus.Errorf("can't unmarshal to Certificate, body:%s, err:%s", resp.Body(), err.Error())
		return out, err
	}
	return out, nil
}

// Endpoints get endpoint from keystone server
func (k *keystone) Endpoints(queries map[string]string) (Endpoints, error) {
	out := Endpoints{}
	req := resty.New().R()
	// set header info for request
	k.setAuth(req).withJsonContentType(req)
	u := k.combine([]string{endpoints}, queries)
	logrus.Debugln("request url is", u.String())
	resp, err := req.Get(u.String())
	if err != nil {
		logrus.Errorf("request[%s] occur exception, err:%s", u.String(), err.Error())
		return out, err
	}

	if err = json.Unmarshal(resp.Body(), &out); err != nil {
		logrus.Errorf("can't unmarshal to Endpoints, body:%s, err:%s", resp.Body(), err.Error())
		return out, err
	}
	return out, nil
}

// Service get service by name from keystone server
func (k *keystone) Service(name string) (Service, error) {
	req := resty.New().R()
	// set header info for request
	k.setAuth(req).withJsonContentType(req)
	u := k.combine([]string{services}, map[string]string{"name": name})
	logrus.Debugln("request url is", u.String())
	resp, err := req.Get(u.String())
	if err != nil {
		logrus.Errorf("request[%s] occur exception, err:%s", u.String(), err.Error())
		return Service{}, err
	}
	out := Services{}
	if err = json.Unmarshal(resp.Body(), &out); err != nil {
		logrus.Errorf("can't unmarshal to Services, body:%s, err:%s", resp.Body(), err.Error())
		return Service{}, err
	}

	if len(out.Services) == 0 {
		return Service{}, errors.New(fmt.Sprintf("not found service[%s]", name))
	}

	return out.Services[0], nil
}

func (k *keystone) combine(paths []string, queries map[string]string) *url.URL {
	return k.combineQueries(k.combinePath(paths...), queries)
}

func (k *keystone) combineQueries(u *url.URL, queries map[string]string) *url.URL {
	if queries == nil {
		return u
	}
	q := u.Query()
	for k, v := range queries {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u
}

func (k *keystone) combinePath(path ...string) *url.URL {
	var err error
	u := &url.URL{}
	if len(path) == 0 {
		u, _ = u.Parse(k.address)
		return u
	}
	// use path[0] instead k.address
	if u, err = url.Parse(path[0]); err == nil && len(u.Scheme) != 0 {
		u.Path = pt.Join(path[1:]...)
		return u
	}
	u, _ = url.Parse(k.address)
	u.Path = pt.Join(u.Path, pt.Join(path...))
	return u
}

func (k *keystone) setAuth(request *resty.Request) *keystone {
	request.SetHeader("X-Auth-Token", k.token)
	request.SetHeader("X-Subject-Token", k.token)
	return k
}

func (k *keystone) withJsonContentType(request *resty.Request) *keystone {
	request.SetHeader("Content-Type", "application/json")
	return k
}
