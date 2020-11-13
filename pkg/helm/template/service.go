package template

import (
	"strconv"
	"strings"
)

const serviceTemplate = `apiVersion: v1
kind: Service
metadata:
  name: {{ include "<CHARTNAME>.fullname" . }}
  labels:
    {{- include "<CHARTNAME>.labels" . | nindent 4 }}
spec:
  {{- if (index .Values.services <INDEX>).<SERVICENAME>.expose.type }}
  type: {{ (index .Values.services <INDEX>).<SERVICENAME>.expose.type }}
  {{- end }}
  {{- if (index .Values.services <INDEX>).<SERVICENAME>.expose.ports }}
  ports:
  {{- (index .Values.services <INDEX>).<SERVICENAME>.expose.ports | toYaml | nindent 2 }}
  {{- end }}
  selector:
    {{- include "<CHARTNAME>.selectorLabels" . | nindent 4 }}`

func (r *render) service(index int, service Service) string {
	o := strings.ReplaceAll(serviceTemplate, CHARTNAME, r.values.ChartName)
	o = strings.ReplaceAll(o, SERVICENAME, service.Name)
	o = strings.ReplaceAll(o, INDEX, strconv.FormatInt(int64(index), 10))
	return o
}
