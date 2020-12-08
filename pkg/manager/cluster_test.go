package manager

import (
	"testing"

	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var (
	clusterUUID = "7d314043-1eb2-43f4-93d3-730091654491"
	keystone    = "http://10.60.41.127:35357/v3"
	token       = "gAAAAABf0CCeY1aDY_Y7AH9QA-SI52LZt3msKEf_-_0ieUTvmLFTrYeK8Vdj6XwCL9PCRPBSjoGihC4eNzkLeGWyFo0yBIxLzQ6y-Am9TQQPpHjBiRdrlYe5D48c8Hev7_b3Uf2V1O694OrnCBYbpp_VEYbHa8E7kA"
)

func TestClusterManager(t *testing.T) {
	manager := NewClusterManager(keystone)
	cluster, err := manager.Client(clusterUUID, token)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = writeKubeconfigFile(cluster); err != nil {
		t.Fatal(err)
	}
}

var (
	globalKubeCfg      = clientcmdapi.NewConfig()
	globalKubeconfFile = "globalkube.conf.yaml"
)

func writeKubeconfigFile(cpy Cluster) (bool, error) {
	clusterName := cpy.Name + "-cluster"
	authName := cpy.Name + "-auth"

	if _, ok := globalKubeCfg.Contexts[cpy.Name]; ok &&
		cpy.ApiAddress == globalKubeCfg.Clusters[clusterName].Server &&
		cpy.CAData == string(globalKubeCfg.Clusters[clusterName].CertificateAuthorityData) &&
		cpy.CertData == string(globalKubeCfg.AuthInfos[authName].ClientCertificateData) &&
		cpy.KeyData == string(globalKubeCfg.AuthInfos[authName].ClientKeyData) {
		return false, nil
	}
	globalKubeCfg.Clusters[clusterName] = clientcmdapi.NewCluster()
	globalKubeCfg.Clusters[clusterName].Server = cpy.ApiAddress
	globalKubeCfg.Clusters[clusterName].CertificateAuthorityData = []byte(cpy.CAData)

	globalKubeCfg.AuthInfos[authName] = clientcmdapi.NewAuthInfo()
	globalKubeCfg.AuthInfos[authName].ClientCertificateData = []byte(cpy.CertData)
	globalKubeCfg.AuthInfos[authName].ClientKeyData = []byte(cpy.KeyData)

	globalKubeCfg.Contexts[cpy.Name] = clientcmdapi.NewContext()
	globalKubeCfg.Contexts[cpy.Name].Cluster = clusterName
	globalKubeCfg.Contexts[cpy.Name].AuthInfo = authName
	globalKubeCfg.CurrentContext = clusterName
	return true, clientcmd.WriteToFile(*globalKubeCfg, globalKubeconfFile)
}
