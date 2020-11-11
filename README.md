当前仓库针对已经配置好的repository进行设计
## 设计思路
### repository设计

#### 同步repository  (已完成)
1. 读取配置,获取repository配置信息
2. 请求repository index.yaml文件并生成repoName-index.yaml

#### chart相关操作
- 获取chart列表 (已完成)
1. 传入参数 {repoName} e.g. bitnami
2. 基于传入参数过滤, 读取repoName-index.yaml遍历所有的chart, 去重chart name相同的chart

- 获取指定chart的所有版本 (已完成)
1. 传入参数 {repoName}/{chartName} e.g. bitnami/apache
2. 通过参数过滤repoName-index.yaml中所有满足条件的chart记录

- 获取chart的所有基本信息
下载chart.tgz, 缓存到指定目录 e.g. {repoName}/{chart.tgz}

1. 获取readme信息 读取readme.md文件
```
[]string{"readme.md", "readme.txt", "readme"}
```
2. 获取values信息 读取values.yaml文件
3. 部署指定chart 
获取chart tgz文件路径 
解析chart中的依赖项 调用helm pkg方法
带上租户信息

### release相关操作
- 获取release列表

- 获取release详细
1. 获取基本信息
2. 调用kubernetes获取依赖资源 

### 创建时候指定namespace
没有namespace需要手动创建

### QA
1. 安装时是否需要指定namespace (不指定可以选择强制创建对应的namespace)
2. 获取列表是否需要指定namespace

> 如需提供同步功能可以删除缓存目录中的相关.tgz文件