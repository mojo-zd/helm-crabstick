package template

import (
	"strconv"
	"strings"
)

const deployTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "<CHARTNAME>.fullname" . }}-<SERVICENAME>
  labels:
    {{- include "<CHARTNAME>.labels" . | nindent 4 }}
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      {{- include "<CHARTNAME>.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      labels:
        {{- include "<CHARTNAME>.selectorLabels" . | nindent 8 }}
    spec:
      {{- if (index .Values.services <INDEX>).<SERVICENAME>.imagePullSecrets }}
      imagePullSecrets:
      {{- range (index .Values.services <INDEX>).<SERVICENAME>.imagePullSecrets }}
      - name: {{ .name }}
      {{- end }}
      {{- else if .Values.imagePullSecrets }}
      imagePullSecrets:
      {{- range .Values.imagePullSecrets }}
      - name: {{ .name }}
      {{- end }}
      {{- end }}
      {{- if (index .Values.services <INDEX>).<SERVICENAME>.hostIPC }}
      hostIPC: {{ (index .Values.services <INDEX>).<SERVICENAME>.hostIPC }}
      {{- end }}
      {{- if (index .Values.services <INDEX>).<SERVICENAME>.volumes }}
      volumes: {{- (index .Values.services <INDEX>).<SERVICENAME>.volumes | toYaml | nindent 6 }}
      {{- end }}
      containers:
      - name: {{ .Chart.Name }}
        image: {{ .Values.image }}
`

func (r *render) deploy(index int, service Service) string {
	o := strings.ReplaceAll(deployTemplate, CHARTNAME, r.values.ChartName)
	o = strings.ReplaceAll(o, SERVICENAME, service.Name)
	o = strings.ReplaceAll(o, INDEX, strconv.FormatInt(int64(index), 10))
	return o
}
