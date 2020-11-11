### Prepare
1. create chart
```
helm create foo
```
2. package `foo` chart
```
helm package foo/ --version 0.1
```
3. mkdir for update repository's index.yaml
```
mkdir chart-tars && mv foo-0.1.tgz chart-tars
```

### 如何生成helm chart 
1. 定义ui
2. 解析ui字段并生成helm chart标准目录 
主要包含以下文件: 
Chart.yaml(chart描述文件)、
依赖项(requirements.yaml)、
templates(模板文件)、
charts(依赖子chart)、
values.yaml、
readme.md
3. 通过helm pkg打包上述生成的文件夹 生成对应的.tgz文件
参考`helm package [chart name] --version semver`
4. 上传helm并更新repository.index文件
1) 新repository
参考`helm repo index chart-tars --url https://charts.bitnami.com/bitnami`
> 需要使用同步工具或者手动同步生成的index.yaml、.tgz文件到helm repository

2) 已经存在的repository
执行`helm repo index chart-tars --merge ~/.cache/helm/repository/xx-index.yaml`
会在chart-tars中生成一个index.yaml

通过上述两个步骤把.tgz、index.yaml上传到repository中

### QA
如果repository中已经存在chart 怎么处理