### 生成chart流程
- 生成资源名称规范:
{{chartName}}-{{serviceName}}
- label包含两个部分
1. helm内置label
2. 用户动态定义的label

- values定义
name: chartName
replicaCount: 1
imagePullSecrets:
- secretName
service1Name:
    replicaCount: 1
    imagePullPolicy: Always
    podAnnotations:
    imagePullSecrets:
    - secretName
    resources:
        limit:
            cpu: 1
            memory: 200Mi
        request:
            cpu:500m
            memory: 150Mi
service2Name:
    imagePullPolicy: IfNotPresent
    configMap:
    - name: cm1
    - name: cm2

3. _helpers.tpl定义
1) 通用name定义
{{- define "<CHARTNAME>.fullname" -}}
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

2) 通用label定义
{{- define "<CHARTNAME>.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

3) 默认replicas定义
{{- define "<CHARTNAME>.replica" }}
{{- default 1 .Values.replicaCount }}
{{- end }}

{{/*Selector过滤标签*/}}
{{- define "<CHARTNAME>.selectorLabels" -}}
app.kubernetes.io/name: {{ include "<CHARTNAME>.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*通用label*/}}
{{- define "<CHARTNAME>.labels" -}}
helm.sh/chart: {{ include "<CHARTNAME>.chart" . }}
{{ include "<CHARTNAME>.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}    

4. 生成的deploy.yaml规范
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "<CHARTNAME>.fullname" . }}-<SERVICENAME>
  labels:
    {{- include "<CHARTNAME>.labels" . | nindent 4 }}
    app: {{ <CHARTNAME>.fullname }}-<SERVICENAME>
spec:
  {{- if not .Values.<SERVICENAME>.replicaCount }}
  replicas: {{ include "<CHARTNAME>.replica" }}
  {{- else }}
  replicas: {{ .Values.<SERVICENAME>.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "<CHARTNAME>.selectorLabels" . | nindent 6 }}
      app: {{ <CHARTNAME>.fullname }}-<SERVICENAME>
  template:
    metadata:
      labels:
        {{- include "<CHARTNAME>.selectorLabels" . | nindent 6 }}
        app: {{ <CHARTNAME>.fullname }}-<SERVICENAME>
    {{- if .Values.<SERVICENAME>.podAnnotations }}
      annotations:
      {{ toYaml .Values.<SERVICENAME>.podAnnotations | nindent 8 }}
    {{- end }}
    spec:
    {{- if .Values.<SERVICENAME>.imagePullSecrets }}
    imagePullSecrets:
    {{- range .Values.<SERVICENAME>.imagePullSecrets }}
    - name: {{ . }}
    {{- end }}
    {{- else if .Values.imagePullSecrets }}
    imagePullSecrets:
    {{- range .Values.imagePullSecrets }}
    - name: {{ . }}
    {{- end }}
    {{- end }}
    


1. 解析ui上传对象 把传入对象按照template分组 生成对应的values文件
e.g.
ui上传了两个资源 deploy: demo, deploy: play

通用部分抽取 `_helper.tpl`
```

```
 
生成的`values.yaml`格式应该如下:
```
demo:
    image: xx.demo:v1
    imagePullPolicy: IfNotPresent
    replicaCount: 1
    secretDependency:
    - secret1
    - secret2
play:
    image: xx.play:v2
    imagePullPolicy: Always
    configMapDependency:
    - config1
    - config2
```

2. 对应的模板文件
`deploy-demo.yaml`:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ include "<chartName>.fullname" . }}
  labels:
    {{- include "<chartName>.labels" . | nindent 4 }}
spec:
  {{- if not .Values.<modeName>.autoscaling.enabled }}
  replicas: {{ .Values.<modeName>.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "<chartName>.selectorLabels" . | nindent 6 }}
  template:
    metadata:
      {{- with .Values.<modeName>.podAnnotations }}
      annotations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      labels:
        {{- include "<chartName>.selectorLabels" . | nindent 8 }}
    spec:
      {{- with .Values.imagePullSecrets }}
      imagePullSecrets:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      serviceAccountName: {{ include "mmm.serviceAccountName" . }}
      securityContext:
        {{- toYaml .Values.podSecurityContext | nindent 8 }}
      containers:
        - name: {{ .Chart.Name }}
          securityContext:
            {{- toYaml .Values.securityContext | nindent 12 }}
          image: "{{ .Values.image.repository }}:{{ .Values.image.tag | default .Chart.AppVersion }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          ports:
            - name: http
              containerPort: 80
              protocol: TCP
          livenessProbe:
            httpGet:
              path: /
              port: http
          readinessProbe:
            httpGet:
              path: /
              port: http
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
```


```
apiVersion: apps/v1
kind: Deployment
{{- $service := (index .Values.services <INDEX>) }}
metadata:
  name: {{ include "<CHARTNAME>.fullname" . }}-<SERVICENAME>
  labels:
    {{- include "<CHARTNAME>.labels" . | nindent 4 }}
    app: {{- include "<CHARTNAME>.fullname" . }}-<SERVICENAME>
spec:
  {{- if not $service.replicaCount }}
  replicas: {{ .Values.replicaCount }}1
  {{- else }}
  replicas: {{ $service.replicaCount }}
  {{- end }}
  selector:
    matchLabels:
      {{- include "<CHARTNAME>.selectorLabels" . | nindent 6 }}
      app: {{- include "<CHARTNAME>.fullname" . }}-<SERVICENAME>
  template:
    metadata:
      labels:
        {{- include "<CHARTNAME>.selectorLabels" . | nindent 6 }}
        app: {{- include "<CHARTNAME>.fullname" . }}-<SERVICENAME>
    {{- if $service.podAnnotations }}
    annotations:
    {{ toYaml $service.podAnnotations | nindent 8 }}
    {{- end }}
    spec:
    {{- if $service.imagePullSecrets }}
    imagePullSecrets:
    {{- range $service.<SERVICENAME>.imagePullSecrets }}
    - name: {{ . }}
    {{- end }}
    {{- else if .Values.imagePullSecrets }}
    imagePullSecrets:
    {{- range .Values.imagePullSecrets }}
    - name: {{ . }}
    {{- end }}
    {{- end }}
```