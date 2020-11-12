package template

const deployTemplate = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "<CHARTNAME>.fullname" . }}-<SERVICENAME>
  labels:
    {{- include "<CHARTNAME>.labels" . | nindent 4 }}
spec:
  replicas: {{ .Values.replicaCount }}1
  selector:
    matchLabels:
      {{- include "<CHARTNAME>.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      labels:
        {{- include "<CHARTNAME>.selectorLabels" . | nindent 8 }}
    spec:
      {{- if (index .Values.services <INDEX>).imagePullSecrets }}
      imagePullSecrets:
      {{- range (index .Values.services <INDEX>).<SERVICENAME>.imagePullSecrets }}
      - name: {{ . }}
      {{- end }}
      {{- else if .Values.imagePullSecrets }}
      imagePullSecrets:
      {{- range .Values.imagePullSecrets }}
      - name: {{ . }}
      {{- end }}
      {{- end }}
      {{- if (index .Values.services <INDEX>).<SERVICENAME>.hostIPC }}
      hostIPC: {{ (index .Values.services <INDEX>).<SERVICENAME>.hostIPC }}
      {{- end }}
      {{- if (index .Values.services <INDEX>).<SERVICENAME>.volumes }}
      {{- range (index .Values.services <INDEX>).<SERVICENAME>.volumes }}
      {{- end }}
      {{- end }}
      containers:
      - name: {{ .Chart.Name }}
        image: "{{ .Values.image }}"
`
