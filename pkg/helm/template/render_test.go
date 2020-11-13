package template

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const testValues = `name: chart
replicaCount: 1
image: nginx:lates
imagePullSecrets:
- name: secretName
services:
- demo:
    type: deployments # kubernetes resource type
    expose: # service define
      type: ClusterIP
      ports:
      - port: 8090
        targetPort: 8090
        protocol: TCP
        name: tcp-8090
      - port: 8091
        targetPort: 8091
        protocol: TCP
        name: tcp-8091
    replicaCount: 1
    imagePullPolicy: Always
    podAnnotations: xxx
    hostIPC: true
    volumes:
    - name: test
      hostPath:
        path: /aaa/bbb
    imagePullSecrets:
    - name: secretNameDemo
    resources:
      limit:
        cpu: 1
        memory: 200Mi
      request:
        cpu: 500m
        memory: 150Mi
`

func TestRender(t *testing.T) {
	render, err := NewRender(testValues)
	if err != nil {
		t.Fatal(err)
		return
	}
	out := render.Replacer()
	for file, content := range out {
		write(fmt.Sprintf("demo/templates/%s", file), content)
	}
	write("demo/templates/_helpers.tpl", render.HelperReplacer())
}

func write(name, context string) {
	ioutil.WriteFile(name, []byte(context), 0775)
}
