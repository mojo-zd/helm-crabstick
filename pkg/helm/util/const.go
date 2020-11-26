package util

const FormatStripingYMDHMS = "2006-01-02 15:04:05"

// the label of deploy、service、cm and so on include when deploy it through helm
const SelectorLabelKey = "app.kubernetes.io/instance"

type KubeKind string

const (
	Deploy                  = "Deployment"
	DaemonSet               = "DaemonSet"
	StatefulSet             = "StatefulSet"
	Job                     = "Job"
	CronJob                 = "CronJob"
	Ingress                 = "Ingress"
	ClusterRole             = "ClusterRole"
	ClusterRoleBinding      = "ClusterRoleBinding"
	RoleBinding             = "RoleBinding"
	Role                    = "Role"
	ConfigMap               = "ConfigMap"
	Secret                  = "Secret"
	ServiceAccount          = "ServiceAccount"
	Service                 = "Service"
	PersistentVolumeClaim   = "PersistentVolumeClaim"
	PersistentVolume        = "PersistentVolume"
	ResourceQuota           = "ResourceQuota"
	ReplicaSet              = "ReplicaSet"
	HorizontalPodAutoscaler = "HorizontalPodAutoscaler"
	NetworkPolicy           = "NetworkPolicy"
	PodSecurityPolicy       = "PodSecurityPolicy"
	StorageClass            = "StorageClass"
)

const (
	DefaultPageSize  = 15
	DefaultPageIndex = 1
)
