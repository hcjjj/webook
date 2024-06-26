# 建设中 🔨


## 项目介绍

**开发环境**

IDE🧑‍💻： [GoLand](https://www.jetbrains.com/go/)

OS🪟🐧：[Ubuntu 22.04.3 LTS (WSL2)](https://ubuntu.com/desktop/wsl)

**开发计划**

- [x] 用户登录服务 👤
  - [x] 注册、登录态校验与刷新
  - [x] 保护登录系统
  - [x] 优化登录性能
  - [x] 短信验证码登录
  - [x] 微信扫码登录
  - [x] 长短 Token 与退出
- [x] 接入配置模块 ⚙️
- [x] 接入日志模块 📋️
- [x] 系统监控埋点 📹️
- [x] 文章服务 📃
  - [x] 新建、修改、保存和发布
  - [x] 阅读、点赞和收藏
  - [x] 榜单模型和缓存
  - [x] 分布式任务调度
- [ ] 评论服务 ✍
- [ ] 用户关系 🧩
- [ ] 搜索服务 🔍
- [ ] 即时通讯 💬
- [ ] 单元/集成测试 ✅

**项目结构**

* 参考 [Kratos](https://go-kratos.dev/)、[go-zero](https://go-zero.dev/) 、[Domain-Driven Design](https://zhuanlan.zhihu.com/p/91525839)
* Service - Repository - DAO (Data Access Object) 三层结构 
  * service：领域服务（domain service），一个业务的完整处理过程
  * repository：领域对象的存储，存储数据的抽象
    * dao：数据库操作
  * domain：领域对象
* handler（和HTTP打交道） → service（主要业务逻辑） → repository（数据存储抽象） → dao（数据库操作）

## 技术栈

**第三方库**

* [gin-gonic/gin](https://github.com/gin-gonic/gin) - HTTP web 框架
  * [Middleware](https://github.com/gin-gonic/contrib) - Collection of middlewares created by the community
  * [cors](https://github.com/gin-contrib/cors) -  Official *cross-origin resource sharing* (CORS) gin's middleware
  * [sessions](https://github.com/gin-contrib/sessions) - Gin middleware for session management
* [dlclark/regexp2](https://github.com/dlclark/regexp2) - full-featured 正则表达式
* [go-gorm/gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
  * [go-gorm/mysql](https://github.com/go-gorm/mysql) - GORM mysql driver
* [golang-jwt/jwt](https://github.com/golang-jwt/jwt) - Golang implementation of JSON Web Tokens (JWT)
* [tencentcloud-sdk-go](https://github.com/TencentCloud/tencentcloud-sdk-go) - Tencent Cloud API 3.0 SDK for Golang
  * [腾讯云 SMS](https://console.cloud.tencent.com/smsv2) 个人用户无法使用短信服务 API
* ~~[shansuma](https://gitee.com/shansuma/sms-sdk-master) - 闪速码 SMS 的 API 接口~~
* [wire](https://github.com/google/wire) - Compile-time Dependency Injection for Go
* [ekit](https://github.com/ecodeclub/ekit) - 支持泛型的工具库
* [mock](https://github.com/uber-go/mock) - GoMock is a mocking framework for the Go programming language
* [go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) - Sql mock driver for golang to test database interactions
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [etcd](https://github.com/etcd-io/etcd) - Distributed reliable key-value store for the most critical data of a distributed system
* [zap](https://github.com/uber-go/zap) - Blazing fast, structured, leveled logging in Go
* [mongo-go-driver](https://github.com/mongodb/mongo-go-driver) - The Official Golang driver for MongoDB
* [sarama](https://github.com/IBM/sarama) - Sarama is a Go library for Apache Kafka
* [prometheus](https://github.com/prometheus)/[client_golang](https://github.com/prometheus/client_golang) - Prometheus instrumentation library for Go applications
* [cron](https://github.com/robfig/cron) - a cron library for go 定时任务
* [redis-lock](https://github.com/gotomicro/redis-lock) - 基于 Redis 实现的分布式锁

**相关环境**

* [Node.js](https://nodejs.org/en)
  * 启动前端：在 webook-fe 目录下先 `npm install` 后 `npm run dev`
  * 前端不完善，采用测试驱动开发
* [Docker](https://www.docker.com/)
  * [镜像源](https://yeasy.gitbook.io/docker_practice/install/mirror)（还是挂代理方便）
  * [mysql](https://hub.docker.com/_/mysql) - An open-source relational database management system (RDBMS)
  * [redis](https://hub.docker.com/r/bitnami/redis) - An open-source in-memory storage
  * [etcd](https://hub.docker.com/r/bitnami/etcd) - A distributed key-value store designed to securely store data across a cluster
  * [mongo](https://hub.docker.com/_/mongo) - MongoDB document databases provide high availability and easy scalability
  * [kafka](https://hub.docker.com/r/bitnami/kafka) - Apache Kafka is a distributed streaming platform used for building real-time applications
  * [prometheus](https://hub.docker.com/r/bitnami/prometheus) - The Prometheus monitoring system and time series database
  * *grafana - The open observability platform*
  * *zipkin - A distributed tracing system*
* [kubernates](https://kubernetes.io/)
  * [Kubernetes cluster architecture](https://kubernetes.io/docs/concepts/architecture/)
  * [kubectl](https://kubernetes.io/docs/tasks/tools/) - The Kubernetes command-line tool
  * [HELM](https://helm.sh/) - The package manager for Kubernetes
  * [ingress-nignx](https://github.com/kubernetes/ingress-nginx) - Ingress-NGINX Controller for Kubernetes
* [wrk](https://github.com/wg/wrk) - Modern HTTP benchmarking tool

## 技术要点
**业务功能**

* 用户登录服务
  * 注册、密码加密存储
  * 登录、登录态校验
    * Cookie + Session
    * Session 存储基于 Redis 实现（多实例部署环境）
      * 但是每次请求都要访问 Redis，性能瓶颈问题
      * 换为 JWT（JSON Web Token）机制
        * 这边有个问题需要解决，多实例部署的退出登录功能
    *   刷新登录状态
      * 在登录校验处执行相关逻辑
      * 控制 Session 的有效期
      * 生成一个新的 Token
  * 保护登录系统
    * 限流（限制每个用户每秒最多发送固定数量的请求  ）
      * 基于 Redis 的 IP 限流
    * 增强登录安全
      * 利用 User-Agent 增强安全性  
  * 优化登录性能
  * 短信验证码登录
    * 验证码是一个独立的功能 （登录、修改密码、危险操作的二次验证）
    * 短信服务也是独立的（方便更换供应商）
    * 验证码登录功能 → 验证码功能 → 短信服务（最基础的服务）
    * 发送验证码功能做用户限流（存入所生成验证码到 Redis 的时候检查有效期）
    * 提高可用性：重试机制、客户端限流、failover（轮询，实时检测）
  * 长短 Token 与登出
  * 微信扫码登录（未完成）
  
* 接入配置模块
  * 不同环境读取不同配置文件
  * viper 接入 etcd，实现远程配置中心
  
* 接入日志模块
  * 抽象日志接口并使用 zap 实现
  * 利用 Gin 的 middleware 打印日志
  * 实现 GORM 的日志接口 
  
* 文章服务

  * 新建、修改、保存和发布

    * 测试驱动开发 TDD，专注于某个功能的实现
    * 文章领域中用户的两重身份
    * 新建、修改、保存、发布
    * 发布时制作库和线上库数据的同步问题
    * *Mysql → MongoDB （未做）*
    * *OSS + CDN （未做）*

  * 阅读、点赞、收藏

    * 三者聚合的表设计、索引的设计策略
    * 采用 Redis 的 map 结构缓存三者的总数
    * 使用软删除缓解性能问题
    * 使用 errgroup.Group 并发查询文章内容和相关数据
    * 用 Kafka 改造阅读计数功能，批量处理消息提高性能
    
  * 榜单模型

    * 综合考虑用户的各种行为、时间的衰减特性和权重因子
    * 滑动窗口 + 优先队列，定时计算热榜文章后缓存
    * 分布式定时任务，为了解决不同实例计算结果可能出现偏差
      * 基于 Redis 的分布式锁
      * 基于 MySQL 实现通用的分布式任务调度机制（乐观锁）

* 监控、埋点和告警

  * 利用 Gin middleware 来统计 HTTP 请求  
* 测试：`wrk -t1 -d1m -c2 http://localhost:8080/test/metric`
  * 利用 GORM 的 Plugin 来监控和数据库有关的信息  
  * 使用 Callback 来监控 GROM 执行时间
  * HTTP 接口里面设计的 Code 字段可以考虑用于监控埋点
  * 利用 Redis 的 Hook 功能监控缓存命中率
  * *接入 OpenTelemetry 集成 zipkin（未做）*
  * *prometheus 集成 Grafana 告警（未做）* 

**编程思想**

* 控制反转（Inversion of Control, IoC）
  * 依赖注入（Dependency Injection）
  * 依赖查找、依赖发现（Go 里面没有）
* 面向接口编程
  * 扩展性强
  * 超前设计，最小化实现

**测试**

* 单元测试

  * Table Driven 模式
  * 最起码做到分支覆盖  
  * 注意与时间相关的测试
* 集成测试

  * 至少测完业务层面的主要正常流程和主要异常流程

**第三方服务治理**

* 提高可用性：重试机制、客户端限流、failover（轮询，实时检测）
* 提高安全性，完整的资源申请与审批流程
* 提高可观测性：日志、metrics、tracing，丰富完善的排查手段

# 部署应用

**环境配置**

```shell
# Ubuntu 22.04.3 LTS
# Golang
wget https://golang.google.cn/dl/go1.22.1.linux-amd64.tar.gz
sudo tar xfz go1.22.1.linux-amd64.tar.gz -C /usr/local
sudo vim /etc/profile
# export GOROOT=/usr/local/go
# export GOPATH=$HOME/go
# export GOBIN=$GOPATH/bin
# export PATH=$GOPATH:$GOBIN:$GOROOT/bin:$PATH
source /etc/profile
go version
go env -w GOPROXY="https://goproxy.cn"
go env -w GO111MODULE=on

# Docker
# 

# Kubernetes
#
```

**用 Kubernetes 部署 Web 服务**

交叉编译

```shell
# Windows → Linux
# powershell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o .\build\webook
go build -tags=k8s -o .\build\webook
# Mac → Linux
GOOS=linux GOARCH=amd64 go build -o /build/webook
```

编写 `Dockerfile`

```dockerfile
# 基础镜像
FROM ubuntu:20.04
# 把编译后的打包进这个镜像，放到工作目录 /app
COPY /build/webook /app/webook
WORKDIR /app
# CMD 是执行命令
# 最佳
ENTRYPOINT ["/app/webook"]
```

```shell
# 构建
docker build -t hcjjj/webook:v0.0.1 .
# 删除
docker rmi -f hcjjj/webook:v0.0.1
# 可以将上述命令都写在 Makefile 里面
```

编写 `k8s.yaml` 后

```shell
# 启动 deployment
kubectl apply -f k8s-webook-deployment.yaml
# 查看 
kubectl get deployments
kubectl get pods
# 查看 POD 的日志
kebectl get logs -f webook-5b4c5b9-4g74z
# 启动 services
kubectl apply -f k8s-webook-service.yaml
# 查看
kubectl get services
# 停止
kubectl delete service webook
kubectl delete deployment webook
```

**用 Kubernetes 部署 Mysql**

```shell
# Mysql 持久化
# 启动
kubectl apply -f k8s-mysql-deployment.yaml
kubectl apply -f k8s-mysql-service.yaml
kubectl apply -f k8s-mysql-pv.yaml
kubectl apply -f k8s-mysql-pvc.yaml
# 查看
kubectl get pv
kubectl get pvc
# 停止
kubectl delete service webook-mysql
kubectl delete deployment webook-mysql
kubectl delete pvc webook-mysql-claim
kubectl delete pv webook-mysql-pv
```

**用 Kubernetes 部署 Redis**

```shell
kubectl apply -f k8s-redis-deployment.yaml
kubectl apply -f k8s-redis-service.yaml
kubectl delete service webook-redis
kubectl delete deployment webook-redis
```

**用 Kubernetes 部署 nginx**

```shell
# 本地环境需要修改 host 到 ip 的映射，host 在 k8s-ingress-nginx.yaml 里面
# ❯ ping  hcjjj.webook.com
# PING hcjjj.webook.com (127.0.0.1) 56(84) bytes of data.
# 64 bytes from localhost (127.0.0.1): icmp_seq=1 ttl=64 time=0.028 ms
# 使用 clash for windows 的话，同时需要在 Bypass Domain/IPNet 中添加 

# 安装 ingress-nignx 
helm upgrade --install ingress-nginx ingress-nginx  --repo https://kubernetes.github.io/ingress-nginx  --namespace ingress-nginx --create-namespace
# 查看
kubectl get service --namespace ingress-nginx
# 启动
kubectl apply -f k8s-ingress-nginx.yaml
# 停止
kubectl get ingresses
kubectl delete ingress webook-ingress
kubectl delete namespaces ingress-nginx
```
