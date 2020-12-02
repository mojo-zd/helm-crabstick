package encrypt

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"

	"github.com/sirupsen/logrus"
)

const (
	bits = 2048
	cn   = "cluster-admin"
	o    = "system:masters"
)

const (
	RsaPrivateKeyType = "RSA PRIVATE KEY"
	CertRequestType   = "CERTIFICATE REQUEST"
)

// GenKey generate rsa private key
func GenKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

// X509ToString parse private key、rsa and so on to string
func X509ToString(typ string, in []byte) (string, error) {
	buf := bytes.Buffer{}
	if err := pem.Encode(&buf, &pem.Block{Type: typ, Bytes: in}); err != nil {
		logrus.Errorf("%s to string failed, err:%s", typ, err.Error())
		return "", err
	}
	return buf.String(), nil
}

// GenCSR generate csr with private key
func GenCSR(key *rsa.PrivateKey) ([]byte, error) {
	return x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:         cn,
			Country:            []string{"CN"},
			Province:           []string{"SC"},
			Locality:           []string{"CD"},
			Organization:       []string{o},
			OrganizationalUnit: []string{"QHWK"},
		}, DNSNames: []string{
			"kubernetes",
			"kubernetes.default",
			"kubernetes.default.svc",
			"kubernetes.default.svc.cluster",
			"kubernetes.default.svc.cluster.local"},
	}, key)
}
