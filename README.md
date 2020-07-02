说明
---
### 目的
使用 nacos 的服务发现能力，生成 prometheus 配置文件，用于 prometheus 自动发现监控服务地址

### 实现思路：
配置 prometheus 指定文件夹下的文件自动发现，
启动一个后台服务定期读取 nacos 上的服务实例信息，
生成 prometheus 配置文件，等待 prometheus 自动生效。

### 实现步骤：

#### 1. 配置 prometheus 文件自动发现如下

```
  - job_name: 'micro-service'
    metrics_path: '/actuator/prometheus'
    file_sd_configs:
      - refresh_interval: 1m
        files:
#          - "/etc/prometheus/conf/apps-svc*.yml"
          - "/etc/prometheus/conf/apps-svc*.json"
```

#### 2. 通过 nacos 的 API 读取 nacos 中的 service 列表，
遍历每个 service 下的所有 instance 和 instance 下的 metadata，
组装 prometheus target 文件。


#### 3. 编译服务
```
go build -o nacos-prometheus-discovery main.go
```
交叉编译多平台可执行文件,文件生成在 target/bin 目录下：
```
sh build.sh
```
生成文件列表：
```
> ll target/bin/
> conf
> nacos-prometheus-discovery-darwin-10.6-amd64
> nacos-prometheus-discovery-linux-amd64
> nacos-prometheus-discovery-windows-4.0-386.exe
> nacos-prometheus-discovery-windows-4.0-amd64.exe


```


#### 4. 启动服务

修改 conf/config.json 中的参数

./nacos-prometheus-discovery conf/config.json

#### EOF