
{{/*name定义*/}}
{{- define "ccc.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*版本定义*/}}
{{- define "ccc.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*replica count定义*/}}
{{- define "ccc.replica" }}
{{- default 1 .Values.replicaCount }}
{{- end }}

{{/*Selector过滤标签*/}}
{{- define "ccc.selectorLabels" -}}
app.kubernetes.io/name: {{ include "ccc.fullname" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*通用label*/}}
{{- define "ccc.labels" -}}
helm.sh/chart: {{ include "ccc.chart" . }}
{{ include "ccc.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}
