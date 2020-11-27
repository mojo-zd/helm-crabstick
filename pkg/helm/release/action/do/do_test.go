package do

import (
	"fmt"
	"os"
	"testing"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/types"

	"github.com/mojo-zd/helm-crabstick/pkg/helm/config"
	"k8s.io/client-go/kubernetes"
)

var values = `## Global Docker image parameters
## Please, note that this will override the image parameters, including dependencies, configured to use the global value
## Current available global Docker image parameters: imageRegistry and imagepullSecrets
##
# global:
#   imageRegistry: myRegistryName
#   imagePullSecrets:
#     - myRegistryKeySecretName

## Bitnami Apache image version
## ref: https://hub.docker.com/r/bitnami/apache/tags/
##
image:
  registry: docker.io
  repository: bitnami/apache
  tag: 2.4.46-debian-10-r82
  ## Specify a imagePullPolicy
  ## ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images
  ##
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ##
  # pullSecrets:
  #   - myRegistryKeySecretName

  ## Set to true if you would like to see extra information on logs
  ## ref:  https://github.com/bitnami/minideb-extras/#turn-on-bash-debugging
  ##
  debug: false

## Bitnami Git image version
## ref: https://hub.docker.com/r/bitnami/git/tags/
##
git:
  registry: docker.io
  repository: bitnami/git
  tag: 2.29.2-debian-10-r10
  pullPolicy: IfNotPresent
  ## Optionally specify an array of imagePullSecrets.
  ## Secrets must be manually created in the namespace.
  ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
  ##
  # pullSecrets:
  #   - myRegistryKeySecretName

## String to partially override apache.fullname template (will maintain the release name)
##
# nameOverride:

## String to fully override apache.fullname template
##
# fullnameOverride:

## Number of Apache replicas to deploy
##
replicaCount: 1

## Pod affinity preset
## ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
## Allowed values: soft, hard
##
podAffinityPreset: ""

## Pod anti-affinity preset
## Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity
## Allowed values: soft, hard
##
podAntiAffinityPreset: soft

## Node affinity preset
## Ref: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity
## Allowed values: soft, hard
##
nodeAffinityPreset:
  ## Node affinity type
  ## Allowed values: soft, hard
  type: ""
  ## Node label key to match
  ## E.g.
  ## key: "kubernetes.io/e2e-az-name"
  ##
  key: ""
  ## Node label values to match
  ## E.g.
  ## values:
  ##   - e2e-az1
  ##   - e2e-az2
  ##
  values: []

## Affinity for pod assignment
## Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
## Note: podAffinityPreset, podAntiAffinityPreset, and  nodeAffinityPreset will be ignored when it's set
##
affinity: {}

## Node labels for pod assignment
## Ref: https://kubernetes.io/docs/user-guide/node-selection/
##
nodeSelector: {}

## Tolerations for pod assignment
## Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
##
tolerations: []

## Get the server static content from a git repository
##
cloneHtdocsFromGit:
  enabled: false
  # repository:
  # branch:
  interval: 60
  resources: {}

## Name of a config map with the server static content
##
# htdocsConfigMap:

## Name of a PVC with the server static content
##
# htdocsPVC:

## Name of a config map with the virtual hosts content
##
# vhostsConfigMap:

## Name of a config map with the httpd.conf file contents
##
# httpdConfConfigMap:

## Pod annotations
## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
##
podAnnotations: {}

## Apache pods' resource requests and limits
## ref: http://kubernetes.io/docs/user-guide/compute-resources/
##
resources:
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  limits: {}
  #   cpu: 100m
  #   memory: 128Mi
  requests: {}
  #   cpu: 100m
  #   memory: 128Mi

## Apache container's liveness and readiness probes
## ref: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
##
livenessProbe:
  enabled: true
  path: "/"
  port: http
  initialDelaySeconds: 180
  periodSeconds: 20
  timeoutSeconds: 5
  failureThreshold: 6
  successThreshold: 1
readinessProbe:
  enabled: true
  path: "/"
  port: http
  initialDelaySeconds: 30
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 6
  successThreshold: 1

## Ingress paramaters
##
ingress:
  ## Set to true to enable ingress record generation
  ##
  enabled: false

  ## Set this to true in order to add the corresponding annotations for cert-manager
  ##
  certManager: false

  ## When the ingress is enabled, a host pointing to this will be created
  ##
  hostname: example.local

  ## Ingress annotations done as key:value pairs
  ## For a full list of possible ingress annotations, please see
  ## ref: https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/annotations.md
  ##
  ## If tls is set to true, annotation ingress.kubernetes.io/secure-backends: "true" will automatically be set
  ## If certManager is set to true, annotation kubernetes.io/tls-acme: "true" will automatically be set
  annotations: {}
  #  kubernetes.io/ingress.class: nginx

  ## The list of additional hostnames to be covered with this ingress record.
  ## Most likely the hostname above will be enough, but in the event more hosts are needed, this is an array
  ## hosts:
  ## - name: example.local
  ##   path: /

  ## The tls configuration for the ingress
  ## ref: https://kubernetes.io/docs/concepts/services-networking/ingress/#tls
  ##
  tls:
    - hosts:
        - example.local
      secretName: example.local-tls

  secrets:
  ## If you're providing your own certificates, please use this to add the certificates as secrets
  ## key and certificate should start with -----BEGIN CERTIFICATE----- or
  ## -----BEGIN RSA PRIVATE KEY-----
  ##
  ## name should line up with a tlsSecret set further up
  ## If you're using cert-manager, this is unneeded, as it will create the secret for you if it is not set
  ##
  ## It is also possible to create and manage the certificates outside of this helm chart
  ## Please see README.md for more information
  # - name: apache.local-tls
  #   key:
  #   certificate:

## Prometheus Exporter / Metrics
##
metrics:
  enabled: false
  ## Bitnami Apache Prometheus Exporter image
  ## ref: https://hub.docker.com/r/bitnami/apache-exporter/tags/
  ##
  image:
    registry: docker.io
    repository: bitnami/apache-exporter
    tag: 0.8.0-debian-10-r207
    pullPolicy: IfNotPresent
    ## Optionally specify an array of imagePullSecrets.
    ## Secrets must be manually created in the namespace.
    ## ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
    ##
    # pullSecrets:
    #   - myRegistryKeySecretName
  ## Metrics exporter pod Annotation and Labels
  ## ref: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
  ##
  podAnnotations:
    prometheus.io/scrape: "true"
    prometheus.io/port: "9117"
  ## Apache Prometheus exporter resource requests and limits
  ## ref: http://kubernetes.io/docs/user-guide/compute-resources/
  ##
  resources:
    # We usually recommend not to specify default resources and to leave this as a conscious
    # choice for the user. This also increases chances charts run on environments with little
    # resources, such as Minikube. If you do want to specify resources, uncomment the following
    # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
    limits: {}
    #   cpu: 100m
    #   memory: 128Mi
    requests: {}
    #   cpu: 100m
    #   memory: 128Mi

## Array to add extra volumes (evaluated as a template)
##
extraVolumes: []

## Array to add extra mounts (normally used with extraVolumes, evaluated as a template)
##
extraVolumeMounts: []

## An array to add extra env vars
##
extraEnvVars: []

## Service paramaters
##
service:
  ## Service type
  ##
  type: LoadBalancer
  ## HTTP Port
  ##
  port: 80
  ## HTTPS Port
  ##
  httpsPort: 443
  ## Specify the nodePort(s) value(s) for the LoadBalancer and NodePort service types.
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
  ##
  nodePorts:
    http: ""
    https: ""
  ## Set the LoadBalancer service type to internal only.
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#internal-load-balancer
  ##
  # loadBalancerIP:
  ## Provide any additional annotations which may be required. This can be used to
  ## set the LoadBalancer service type to internal only.
  ## ref: https://kubernetes.io/docs/concepts/services-networking/service/#internal-load-balancer
  ##
  annotations: {}

  ## Enable client source IP preservation
  ## ref http://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#preserving-the-client-source-ip
  ##
  externalTrafficPolicy: Cluster
`

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
	namespace   = "demo"
	releaseName = "bn"
	chartName   = "apache"
	version     = "8.0.0"
)

func TestUpgrade(t *testing.T) {
	_, err := NewDoer(getClient(t), conf).Upgrade(releaseName, chartName, version, values, namespace)
	if err != nil {
		t.Fatal(err)
	}
}

func TestInstall(t *testing.T) {
	doer := NewDoer(getClient(t), conf)
	rls, err := doer.Install(types.ReleaseCreateOptions{
		ChartName: chartName,
		Name:      releaseName,
		Namespace: namespace,
		Values:    values,
		Options:   types.DoerOptions{Annotation: map[string]string{"author": "mojo"}},
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(rls.Name, rls.Chart.Metadata.Annotations["author"])
}

func TestDelete(t *testing.T) {
	doer := NewDoer(getClient(t), conf)
	if _, err := doer.Delete(releaseName, namespace); err != nil {
		t.Error("delete release failed", err)
		return
	}
	t.Log("delete success")
}

func getClient(t *testing.T) kubernetes.Interface {
	restConf, err := conf.ConfigFlags().ToRESTConfig()
	if err != nil {
		t.Fatal(err)
		return nil
	}
	client, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		t.Fatal(err)
	}

	return client
}
