### 利用harbor作为helm chart server
#### 安装harbor并开启helm chart server功能
harbor.yaml请自行配置
```
sh install.sh --with-chartmuseum --with-clair
```

#### 自制chart并上传到harbor helm repository
```
helm create demo 
helm package demo --version 0.1
```
通过harbor项目下的`Helm Chart`tab点击上传.tgz文件

#### 添加harbor helm repository
```
helm repo add --username=xxx --password=xxx demo http://x.x.x.x:xxx/chartrepo
```
