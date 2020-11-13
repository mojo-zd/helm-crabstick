package get

import (
	"regexp"
	"testing"
)

const manifest = `        # Source: apache/templates/svc.yaml
        apiVersion: v1
        kind: Service
		kind: Service
        metadata:
          name: mn-apache
          labels:
            app.kubernetes.io/name: apache
            helm.sh/chart: apache-7.6.0
            app.kubernetes.io/instance: mn
            app.kubernetes.io/managed-by: Helm
          annotations: 
            {}
        spec:
          type: LoadBalancer
          externalTrafficPolicy: "Cluster"
          ports:
            - name: http
              port: 80
              targetPort: http
            - name: https
              port: 443
              targetPort: https
          selector:
            app.kubernetes.io/name: apache
            app.kubernetes.io/instance: mn
        ---
        # Source: apache/templates/deployment.yaml
        apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: mn-apache
          labels:
            app.kubernetes.io/name: apache
            helm.sh/chart: apache-7.6.0
            app.kubernetes.io/instance: mn
            app.kubernetes.io/managed-by: Helm
        spec:
          selector:
            matchLabels:
              app.kubernetes.io/name: apache
              app.kubernetes.io/instance: mn
          replicas: 1
          template:
            metadata:
              labels:
                app.kubernetes.io/name: apache
                helm.sh/chart: apache-7.6.0
                app.kubernetes.io/instance: mn
                app.kubernetes.io/managed-by: Helm
            spec:
              
              hostAliases:
                - ip: "127.0.0.1"
                  hostnames:
                    - "status.localhost"
              affinity:
                podAffinity:
                  
                podAntiAffinity:
                  preferredDuringSchedulingIgnoredDuringExecution:
                    - podAffinityTerm:
                        labelSelector:
                          matchLabels:
                            app.kubernetes.io/name: apache
                            app.kubernetes.io/instance: mn
                        namespaces:
                          - aaa
                        topologyKey: kubernetes.io/hostname
                      weight: 1
                nodeAffinity:
                  
              containers:
                - name: apache
                  image: docker.io/bitnami/apache:2.4.46-debian-10-r62
                  imagePullPolicy: "IfNotPresent"
                  env:
                    - name: BITNAMI_DEBUG
                      value: "false"
                  ports:
                    - name: http
                      containerPort: 8080
                    - name: https
                      containerPort: 8443
                  livenessProbe:
                    httpGet:
                      path: /
                      port: http
                    initialDelaySeconds: 180
                    periodSeconds: 20
                    timeoutSeconds: 5
                    successThreshold: 1
                    failureThreshold: 6
                  readinessProbe:
                    httpGet:
                      path: /
                      port: http
                    initialDelaySeconds: 30
                    periodSeconds: 10
                    timeoutSeconds: 5
                    successThreshold: 1
                    failureThreshold: 6
                  resources:
                    limits: {}
                    requests: {}
                  volumeMounts:
              volumes:
`

func TestManifest(t *testing.T) {
	rr := `kind:\s(\w+)`
	gg, err := regexp.Compile(rr)
	if err != nil {
		t.Error(err)
		return
	}
	kk := gg.FindAllString(manifest, -1)
	single := make(map[string]bool)
	for _, k := range kk {
		if _, ok := single[k]; ok {
			continue
		}
		single[k] = true
		t.Log("kkkk:", k)
	}
}
