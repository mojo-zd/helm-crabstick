### Helm工作流程
#### 添加仓库 
1. 缓存仓库信息到$home/.config/helm、缓存仓库index文件到$home/.cache/helm/repository目录

#### 获取chart
1. 从repository-index.yaml中遍历chart信息

#### 获取chart详细信息
e.g. 获取manifest、readme、notes等等
1. 从repository index.yaml中获取chart .tgz下载地址并下载到本地
2. 解压.tgz文件 遍历解压后的文件内容

#### 安装chart
安装chart以后helm需要记录一条对应的信息来记录chart部署后的状态, 即为release对象
1. 获取chart对象,检查依赖项并解压相关.tgz文件
2. 渲染模板, 存储release信息到storage.Driver
3. 下发渲染后的文件

#### 获取release
从storage.Driver中获取相关release信息

> helm storage如果不指定 默认存储在secrets中, 目前支持memory、secrets、configmap

### 如何集成helm
#### 系统内置repository
从内置repository查找相关chart
#### 安装chart
客户端传入tenant信息, 服务器端获取chart对象并写入tenant信息到annotation中。如下为chart的metadata对象
```
type Metadata struct {
	// The name of the chart
	Name string `json:"name,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	Home string `json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	Sources []string `json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	Version string `json:"version,omitempty"`
	// A one-sentence description of the chart
	Description string `json:"description,omitempty"`
	// A list of string keywords
	Keywords []string `json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*Maintainer `json:"maintainers,omitempty"`
	// The URL to an icon file.
	Icon string `json:"icon,omitempty"`
	// The API Version of this chart.
	APIVersion string `json:"apiVersion,omitempty"`
	// The condition to check to enable chart
	Condition string `json:"condition,omitempty"`
	// The tags to check to enable chart
	Tags string `json:"tags,omitempty"`
	// The version of the application enclosed inside of this chart.
	AppVersion string `json:"appVersion,omitempty"`
	// Whether or not this chart is deprecated
	Deprecated bool `json:"deprecated,omitempty"`
	// Annotations are additional mappings uninterpreted by Helm,
	// made available for inspection by other applications.
	Annotations map[string]string `json:"annotations,omitempty"`
	// KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	KubeVersion string `json:"kubeVersion,omitempty"`
	// Dependencies are a list of dependencies for a chart.
	Dependencies []*Dependency `json:"dependencies,omitempty"`
	// Specifies the chart type: application or library
	Type string `json:"type,omitempty"`
}
```