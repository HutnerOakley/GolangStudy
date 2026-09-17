# GolangStudy

一个用于学习和练习 Go 语言的项目。仓库包含 Go 基础语法示例，以及一个基于 TCP 和 goroutine 实现的简易命令行聊天室。

## 学习内容

| 目录 | 内容 |
| --- | --- |
| `InitLib1/`、`InitLib2/` | 包初始化与 `init` 函数 |
| `defers/` | `defer`、文件读取、锁和 HTTP Handler 示例 |
| `funcs/` | 函数、多返回值和指针参数 |
| `generics/` | 泛型函数与类型约束 |
| `gins/` | Web 响应结构体与状态码类型示例 |
| `gotoutines/` | goroutine、channel 和并发示例 |
| `iotas/` | `iota` 与枚举类型 |
| `oops/` | 结构体、方法、接口和空接口值 |
| `recovers/` | `panic` 与 `recover` |
| `reflects/` | 反射基础 |
| `slices/` | slice 和 map 操作 |
| `var/` | 变量与常量 |
| `servers/` | TCP 聊天室的服务端、客户端和用户逻辑 |
| `cmd/client.go` | 聊天室客户端入口 |
| `main.go` | 示例代码及聊天室服务端入口 |

## TCP 聊天室

聊天室默认监听 `127.0.0.1:8888`，支持：

- 用户上线、下线广播
- 公聊消息广播
- 查询在线用户
- 修改用户名
- 用户之间私聊
- 连接长时间不活跃时自动断开

### 启动服务端

在项目根目录打开一个终端：

```powershell
go run .
```

根入口会先运行部分 Go 学习示例，然后启动 TCP 服务端。

### 启动客户端

保持服务端运行，在另一个终端进入同一项目目录：

```powershell
go run ./cmd
```

客户端连接成功后会显示以下菜单：

```text
1.公聊模式
2.私聊模式
3.更新用户名
0.退出
```

可以同时启动多个客户端，用于测试广播、在线用户列表和私聊功能。

如果当前位于项目的上一级目录，也可以使用：

```powershell
go -C GolangStudy run ./cmd
```

## 环境要求

- Go 1.25 或更高版本

安装依赖：

```powershell
go mod download
```

## 编译检查

项目目前没有自动化测试，可以使用下面的命令编译检查所有包：

```powershell
go test ./...
```

## 当前说明

这是一个学习用途的项目。服务端和客户端当前使用固定地址 `127.0.0.1:8888`，暂未包含消息持久化、身份认证或加密传输等生产环境功能。
