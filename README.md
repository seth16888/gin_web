# Gin Web 服务基础项目
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-blue?style=flat-square)
![Go](https://img.shields.io/badge/Go-1.20%2B-green?style=flat-square)

## 简介
本项目是一个使用 Gin 框架开发的基础 Web 服务项目。Gin 是一个用 Go 语言编写的轻量级 Web 框架，以高性能、简洁易用等特点著称。此项目提供了一个基本的 Web 服务结构，可作为开发更复杂 Web 应用的起点。

## 功能特性
- **快速启动**：基于 Gin 框架，能快速搭建 Web 服务，节省开发时间。
- **路由管理**：提供简单的路由配置示例，方便开发者根据需求定制路由。
- **中间件支持**：支持使用 Gin 中间件进行请求处理，增强服务的功能和可扩展性。

## 项目结构

以下是本项目的主要目录结构及其说明：

```plaintext
root/
├── conf/               # 配置文件
├── docs/               # 项目文档
├── internal/           # 内部业务逻辑
│   ├── bootstrap/      # 启动初始化
│   ├── cmd/            # 命令行工具
│   ├── config/         # 配置管理
│   ├── di/             # 依赖注入
│   ├── server/         # 服务器相关代码
│   │   ├── router/     # 路由配置
│   │   ├── middleware/ # 中间件
│   │   └── handler/    # 处理请求的逻辑
│   │   └── server.go   # 服务器启动
├── pkg/                # 公共工具包
├── tests/              # 测试文件
├── scripts/            # 脚本文件
├── go.mod              # Go 模块依赖
├── go.sum              # Go 模块校验
└── README.md           # 项目说明文档
```

### 目录说明
- **`cmd`**：包含项目的入口文件，程序从此处启动。
- **`internal`**：存放项目的核心业务逻辑，按照功能模块划分目录。
- **`conf`**：项目的配置文件，可根据不同环境进行配置。
- **`docs`**：项目的相关文档，如设计文档、接口文档等。
- **`tests`**：单元测试和集成测试文件，确保代码质量。

## 快速开始

### 环境准备
确保你已经安装了 Go 语言环境（版本 1.20 及以上）。

### 克隆项目
```bash
git clone https://gitee.com/seth16888/gin_web.git
cd gin_web
```

### 安装依赖
```bash
go mod tidy
```
### 配置项目
编辑 config/config.yaml 文件，根据需要修改配置参数。

### 运行项目
```bash
.\scripts\build.cmd
.\bin\gin_web.exe -c conf\config.yaml
```

### 访问服务
启动成功后，你可以通过 http://localhost:10010/ping 访问服务。

### 路由配置
路由配置位于 internal/server/router/router.go 文件中，你可以根据需要添加或修改路由。

```go
func registerRoutes(r *gin.Engine) {
	r.GET("/ping", handler.NewHealthHandler().Ping)
	r.GET("/health", handler.NewHealthHandler().Health)
}
```

## 贡献
如果你想为这个项目做出贡献，请遵循以下步骤：

1. Fork 本仓库。
2. 创建一个新的分支： git checkout -b feature/your-feature 。
3. 提交你的更改： git commit -m 'Add some feature' 。
4. 推送至远程分支： git push origin feature/your-feature 。
5. 提交 Pull Request。
