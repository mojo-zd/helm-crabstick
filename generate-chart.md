### 生成chart流程
1. 解析ui上传对象 把传入对象按照template分组 生成对应的values文件
e.g.
ui上传了两个资源 deploy: demo, deploy: play 
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
`deploy-play.yaml`:
```

```
