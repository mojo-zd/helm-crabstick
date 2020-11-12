package template

import (
	"io/ioutil"
	"testing"
)

const testValues = `
name: ccc
replicaCount: 1
imagePullSecrets:
- secretName
services:
- demo:
    type: deployments
    replicaCount: 1
    imagePullPolicy: Always
    podAnnotations: xxx
    imagePullSecrets:
    - name: secretName
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
	deploy := render.Replacer()
	write(deploy[0], "demo/templates/deploy.yaml")
	write(render.HelperReplacer(), "demo/templates/_helpers.tpl")
}

func write(context, name string) {
	ioutil.WriteFile(name, []byte(context), 0775)
}
