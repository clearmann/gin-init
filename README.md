## 项目模板介绍
gin-init 是一个基于 Gin 框架的项目初始化模板，它是由Golang生态中 **最新** + **最火** 技术整合而成的，它们的组合可以帮助你快速构建一个高效、可靠的应用程序。该脚手架包括配置管理、日志记录、中间件、依赖注入、数据库集成和测试等常见功能，来帮助开发者更快速地启动项目。  

## 模板特点：
- - -
### 主流技术 & 特性
- Docker
- Gin
- GORM
- Viper: 支持 TOML (默认)、YAML 等常用格式作为配置文件
- Casbin
- Zap
- Wire
- Make
### 数据存储
- MySQL 数据库
- Redis 内存数据库
- Elasticsearch 搜索引擎
- 腾讯云 COS 对象存储
### 业务特性
- Redis 分布式登录
- 全局请求响应拦截器(记录日志)
- 全局异常处理器
- 自定义错误码
- 封装通用响应类
- Swagger 接口文档
- 全局跨域处理
- 多环境配置
- - - -
## 业务功能
- 用户登录、注册、更新
- jwt 权限验证
- 全量同步ES、增量同步 ES 定时任务
- 支持微信扫码登录
- 支持手机验证码登录
- 支持分业务的文件上传
- 更多功能还在梳理中...
### 单元测试
- 示例单元测试类
### 架构设计
- 合理分层
## 快速上手
> 所有需要修改的地方加上 TODO(未完成)
### 目录结构
```
.
├── api
│   └── v1
│── bin                   -- 生成的二进制文件
├── cmd                   -- 应用程序的入口，包含了不同的子命令
│   ├── migration
│   ├── server
│   │   ├── wire
│   │   │   ├── wire.go      -- Wire配置文件，定义了server子命令需要的依赖关系
│   │   │   └── wire_gen.go  -- 自动编译生成的Wire文件
│   │   └── main.go
│   └── task
├── config                -- 配置文件
├── deploy                -- 部署相关的文件，如 Dockerfile 、 docker-compose.yml等
├── docs
├── internal              -- 应用程序的主要代码，按照分层架构进行组织
│   ├── handler           -- 处理 HTTP 请求，调用业务逻辑层的服务，返回 HTTP 响应
│   ├── middleware
│   ├── model             -- 数据模型，定义了业务逻辑层需要的数据结构
│   ├── repository        -- 数据访问对象，封装了数据库操作，提供了对数据的增删改查
│   ├── server            -- HTTP 服务器，启动 HTTP 服务，监听端口，处理 HTTP 请求
│   └── service           -- 服务，实现具体的业务逻辑，调用数据访问层repository
├── pkg                   -- 公共的代码，包括配置、日志、HTTP 等
│   ├── config            -- 配置文件的读取和解析
│   ├── jwt               -- 数据模型，定义了业务逻辑层需要的数据结构
│   ├── log               -- 日志相关的代码，如日志的初始化、日志的写入等
│   └── more...
├── script                -- 脚本文件，用于部署和其他自动化任务
├── storage               -- 存储文件，如日志文件
├── test                  -- 测试代码
│   ├── mocks
│   └── server
├── Makefile              -- Makefile，用于构建、测试、打包等 
├── go.mod
└── go.sum
```

### 启动项目
您可以使用以下命令快速启动项目：
```
make build
```
此命令将启动 Golang 脚手架项目。
### 编译 wire.go 
您可以使用以下命令快速编译 wire.go：
```
cd cmd/server && wire
```
此命令会自动寻找项目中的wire.go文件，并快速编译生成所需的依赖项。
### 配置文件
#### 指定配置文件来启动
gin-init 使用 Viper 库来管理配置文件。
默认会加载```config/local.yml```，你可以使用环境变量或参数来指定配置文件路径
```
set APP_CONF=config\prod.yml && go run ./cmd/server
```
或者使用传参的方式:```go run ./cmd/server -conf=config/prod.yml```

#### 读取配置项
您可以在 ```config``` 目录下创建一个名为 ```local.yaml``` 的文件来存储您的配置信息。例如：
```yaml
data:
  mysql:
    user: root:123456@tcp(127.0.0.1:3380)/user?charset=utf8mb4&parseTime=True&loc=Local
  redis:
    addr: 127.0.0.1:3306
    password: ""
    db: 0
    read_timeout: 0.2s
    write_timeout: 0.2s
```
您可以在代码中使用依赖注入```conf *viper.Viper```来读取配置信息。

tips：通过参数进行依赖注入之后，别忘记执行 wire命令生成依赖文件。
## 日志
使用 Zap 库来管理日志。您可以在 config 中配置日志。例如：
```
log:
  log_level: info
  encoding: json               # json or console
  log_file_name: "./storage/logs/server.log"
  max_backups: 30              # 日志文件最多保存多少个备份
  max_age: 7                   #  文件最多保存多少天
  max_size: 1024               #  每个日志文件保存的最大尺寸 单位：M
  compress: true               # 是否压缩
```
您可以在代码中使用以下方式来记录日志：
```
TODO
```
## 数据库
使用 GORM 库来管理数据库。您可以在 config 目录下配置数据库。例如：
```yaml
data:
  mysql:
    user: root:123456@tcp(127.0.0.1:3380)/user?charset=utf8mb4&parseTime=True&loc=Local
  redis:
    addr: 127.0.0.1:3306
    password: ""
    db: 0
    read_timeout: 0.2s
    write_timeout: 0.2s
```
您可以在代码中使用以下方式来连接数据库：
```
TODO
```
需要注意的是项目中的xxxRepository、xxxService、xxxHandler都是基于interface实现，
这就是所谓的面向接口编程，它可以提高代码的灵活性、可扩展性、可测试性和可维护性，是Go语言非常推崇的一种编程风格。
比如上面的代码我们写成了
```
type UserRepository interface {
	FirstById(id int64) (*model.User, error)
}
type userRepository struct {
	*Repository
}
```
而不是直接写成
```
type UserRepository struct {
	*Repository
}
```
单元测试是基于interface特性进行mock操作的。
## 测试
该项目使用 testify、redismock、gomock、go-sqlmock等 库来编写测试。

您可以使用以下命令运行测试：
```
go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./.gin-init/coverage.out ./test/server/...
go tool cover -html=./.gin-init/coverage.out -o coverage.html
```
上面的命令将会生成一个html文件coverage.html，我们可以直接使用浏览器打开它，然后我们就会看到详细的单元测试覆盖率。

## 结论
gin-init 是一个非常实用的 Golang 应用脚手架，它可以帮助您快速构建高效、可靠的应用程序。希望本指南能够帮助您有所学习，并且您可以根据自己的需求进行扩展和定制，请不要忘记给本项目一个star。
