package encrypt

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestX509(t *testing.T) {
	key, err := GenKey()
	if err != nil {
		t.Fatal("generate private key failed", err)
	}
	csr, err := GenCSR(key)
	if err != nil {
		t.Fatal("generate csr failed", err)
	}
	buf := bytes.Buffer{}
	pem.Encode(&buf, &pem.Block{Type: CertRequestType, Bytes: csr})
	t.Log("csr:", buf.String())
	keyBuf := bytes.Buffer{}
	pem.Encode(&keyBuf, &pem.Block{Type: RsaPrivateKeyType, Bytes: x509.MarshalPKCS1PrivateKey(key)})
	t.Log("private key:", keyBuf.String())
}
