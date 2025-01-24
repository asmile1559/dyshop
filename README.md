# dyshop

# Overview

## Content
- [How to Run](#how-to-run-thinking_face)
  - [Start a Microservice App](#start-a-microservice-app)
  - [Testing Microservice](#testing-microservice)
- [How to Collaborate](#how-to-collaborate)
- [Project Architecture](#project-architecture)
- [Framework](#framework)
- [Project Quick Start](#project-quick-start)
- [Progress](#Progress)
  - [Framework](#framework-1)
  - [Auth](#auth)

#  How to Run 🤔
## Start a Microservice App
`app`目录下的每一个目录都是一个微服务。微服务的运行方式都是相同的。
1. 进入微服务目录（以auth服务为例）
```shell
dyshop$ cd app/auth
```
2. 运行微服务程序
```shell
dyshop/app/auth$ go run .
```
> <strong>⚠注意：</strong>  
> 各个微服务请从自己的微服务目录中运行，不要在项目路径或其他路径中运行。  
> 微服务的搜索路径与起始路径有关，在其他地方运行服务，可能会导致<strong>`conf`文件读取错误❌</strong>或其他问题。

3. 微服务启动后，默认没有输出，但是会使终端处于阻塞状态。你可以自己控制微服务的输出。

## Testing Microservice
为了测试微服务是否可用，在本项目中有三种进行测试的方式。
1. 离线测试：  
使用`go`的`go test` 编写test函数，进行离线测试。

```go
package service

// this file must end with _test.go
// for example: example_test.go
import (
	"testing"
)

// optional function
// if you have some common initialization operations 
// for all test functions please write in this function
func init(){
	// init operations before test...
}

// TestXXX this function must start with Test
func TestXXX(t *testing.T) {
	// write your test code here...
}
```
运行go test的方式也比较简单
```shell
go test -run TestXXX -v # 运行某个测试函数，并且详细输出结果
go test -file example_test.go -v # 运行某个测试文件的所有测试函数
```
> 💡这种方式能够在本地得到service函数的执行结果，并且可以设置测试样例，进行覆盖测试。  
> 💡这也是一种在`非main包`情况下运行某个函数的方式。
2. 本地RPC调用测试  
我们已经提供了本地的RPC客户端的实现，就在当前微服务的cmd/client目录下。
```go
package main

import (
	"fmt"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	// ...
)

func main() {
    // initialization operations
	cc, err := grpc.NewClient("localhost:"+viper.GetString("server.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}
	cli := pbauth.NewAuthServiceClient(cc)
	// you can call your rpc functions here by grpc
	resp, err := cli.DeliverTokenByRPC(context.TODO(), &pbauth.DeliverTokenReq{UserId: 1})
    // other operations
}
```
执行命令如下：
```shell
# 1. start up microservice server
dyshop/app/auth$ go run . 
# 2. open a new terminal and run client
dyshop/app/auth$ go run cmd/client/main.go
```
> ❗这种方式首先需要启动RPC服务，也即运行微服务本身。

3. 标准请求方式  
标准请求方式是外部请求从浏览器发起，经过frontend（gateway）进行转发处理，并最终将结果返回给浏览器的过程。  
这是实际上我们最终需要实现的。  
执行命令如下：
```shell
# 1. start up microservice server
dyshop/app/auth$ go run . 
# 2. open a new terminal and start up frontend
dyshop/app/frontend$ go run .
# 3. request by browser, postman or curl
# recommend browser and postman
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1}' \
  http://localhost:10166/test/login
```
> 👍 推荐在进行最后功能测试时使用这种方式

# How to collaborate
0. 环境（框架功能的开发环境）,**一般情况下<font color="red">不需要</font>改自己的环境**
> Go version: go1.23.5 linux/amd64  
> IDE：GoLand/VSCode（VSCode 有时候会出现找不到模块的报错）  
> OS：Linux/（Debian12，Ubuntu2204)  

1. 将代码克隆到本地。**推荐使用ssh模式**。也可以fork之后clone自己的。
```shell
git clone git@github.com:asmile1559/dyshop.git
````
2. 新建一个分支
```shell
# example需要替换为自己负责的模块
git branch feat/example
git checkout feat/example
```

3. 安装 `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`.
```shell
# 1. 安装 protoc
# 参考 https://grpc.io/docs/protoc-installation/
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v29.3/protoc-29.3-linux-x86_64.zip
unzip protoc-29.3-linux-x86_64.zip -d $HOME/.local

# 可以写到 $HOME/.profile 或 $HOME/.bashrc 中
export PATH="$PATH:$HOME/.local/bin" 

# 2. 安装 protoc-gen-go 和 protoc-gen-go-grpc
# 参考：https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# 可以写到 $HOME/.profile 或 $HOME/.bashrc 中
export PATH="$PATH:$(go env GOPATH)/bin"
```

4. 生成proto文件，**可以不生成（当需要修改proto文件时，需要运行下面的命令）**

```shell
make gen-frontend-proto
make gen-backend-proto
```

5. 开始编写代码吧。

微服务的业务代码一般在下面三个文件中编写。
- `app/xxx/main.go`
- `app/xxx/handler.go`
- `app/xxx/service/*.go`
 
> 👍 推荐 **（不是硬性规定）** 👍
> - 😀 在 `app/xxx/main.go`文件中编写运行初始化、服务注册等部分代码。
> - 😁 在 `app/xxx/handler.go`中编写与RPC相关部分的代码。
> - 😊 在 `app/xxx/service/*.go`中编写具体的业务代码。  
> 
> 💡 `auth`部分的业务逻辑已经实现，如果需要编写自己部分的业务逻辑，可以参考。

6. 推送分支
```shell
# 追踪文件
git add .
# 本地提交
git commit -m "example commit"
# 如果是多人协调一个分支
git pull origin feat/example
# 本地处理冲突
...
# 推送到远端
git push origin feat/example 
```
7. **发起一个PR请求**  
你可以考虑自己自己与主分支merge，也可以和别人进行code review后进行merge。  
推荐与别人讨论后再merge。

# Project Architecture

```text
dyshop/ # 工程根目录
├── app # 微服务目录
│   ├── auth # 鉴权与认证服务
│   │   ├── biz # 业务代码
│   │   │   ├── dal # 数据库相关操作
│   │   │   ├── model # 模型定义
│   │   │   └── service # 服务代码
│   │   │       ├── deliver_token.go 
│   │   │       ├── deliver_token_test.go
│   │   │       ├── verify.go
│   │   │       └── verify_test.go 
│   │   ├── cmd # 客户端实现
│   │   │   └── client
│   │   │       └── main.go
│   │   ├── conf # 配置文件
│   │   │   ├── config.yaml
│   │   │   ├── model.conf
│   │   │   └── policy.csv
│   │   ├── docker-compose.yaml # 当前服务启动的容器
│   │   ├── go.mod 
│   │   ├── go.sum
│   │   ├── handler.go # RPC相关代码
│   │   ├── main.go # 微服务入口
│   │   ├── middleware # 中间件函数
│   │   ├── script # 脚本
│   │   └── utils # 当前微服务使用的工具函数
│   ├── ... # 其他微服务
├── pb # protoc 生成的文件
│   ├── backend
│   ├── frontend
│   └── go.mod
├── proto # proto源文件
│   ├── backend
│   └── frontend
├── utils 
│    ├── balancerx # 负载均衡
│    ├── configx # 配置
│    ├── ctool # 加密
│    ├── db # 数据库
│    ├── example # 示例
│    ├── filex # 文件操作
│    ├── registryx # 服务注册
│    ├── jwt # token
│    ├── logx # 日志
│    └── go.mod
├── assets # README.md 使用的资源目录
├── deploy # 微服务部署
├── README.md # 说明
├── dyshop.postman_collection.json # postman请求测试文件
├── go.work # workspace文件
└── Makefile # 常用命令
```
    
# Framework:

![Framework](./assets/framework.jpg)

1. 外部请求通过 RESTful API 发送至 gateway(frontend)
2. 当一个app启动时，会向注册中心进行注册
3. 内部服务之间的访问，通过注册中心找到对应的服务，并通过grpc调用传输。
4. 建议每一个微服务使用自己独立的数据库，这样一方面加快数据库的查找速度，另一方面更加安全。

# Project Quick Start
1. 在 proto 目录下的前端和后端模块中，编写所需要的proto文件（或者修改原本的proto文件）。
```protobuf
syntax = "proto3"; // proto 协议版本 [required]

package hello; // 当前proto文件的包名，用于 proto 之间的 import [required]

// option go_package = "example.com/user/project/whatever;whatever"
//   ↑        ↑                      ↑              ↑        ↑
// proto选项 | 生成的 go 包的参数   按照项目填写      生成go包名 |其他包的引用名
// ⚠注意：生成的go包路是 go_out 路径加 example.com/user/project/whatever/**.go
option go_package = "github.com/asmile1559/dyshop/pb/backend/hello;hello";

// The greeter service definition.
service Greeter {
//       ↑
//  service的名字，一个大的service有很多小的rpc调用（小的service）
//  这里的service会由grpc提供 client 和 server 接口
  rpc SayHello (HelloRequest) returns (HelloReply) {}
// ↑     ↑           ↑           ↑         ↑        ↑
//声明rpc|RPC调用名|RPC调用参数|声明返回值|RPC调用返回值|可拓展参数
}

// The request message containing the user's name.
message HelloRequest {
// ↑         ↑
//声明消息   消息名，对应调用参数和返回值
  string name = 1; 
//  ↑     ↑     ↑
// 类型  参数名 序列号
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```

2. 在工程目录下，运行 `make gen-backend-proto` 或者 `make gen-frontend-proto`。它会在pb目录下的对应位置生成你需要的go文件。
3. 在app目录下找到或者新建你的微服务。如 `auth`。并按照project architecture的结构创建文件。
4. 模仿某一个app的`biz/service`的内容编写自己的service，最好添加测试，可以测试自己`service`的可用性。推荐的写法如下：
```go
package service

import (
    pbhello "github.com/asmile1559/dyshop/pb/backend/hello"
    "context"
)

type SayHelloService struct {
	ctx context.Context
}

func NewSayHelloService(c context.Context) *SayHelloService {
	return &SayHelloService{ctx: c}
}

func (s *SayHelloService) Run(req *pbhello.HelloRequest) (*pbhello.HelloReply, error) {
	// TODO: finish your business code...
	//
	return 
}


```

5. 在handler中实现对应rpc的方法。

```go
package main

// 1. 导入对应的依赖
import (
    "context"
	// service 是自己实现的service
    service "github.com/asmile1559/dyshop/app/hello/biz/service"
	// pbhello 是 protoc 生成的的文件
    pbhello "github.com/asmile1559/dyshop/pb/backend/hello"
)

type Greeter struct{
	// 包含这个未实现的 server 即可
	pbhello.UnimplementedGreeterServer
}

func (s *Greeter) SayHello(ctx context.Context, req *pbhello.HelloRequest) (*pbhello.HelloReply, error) {
	// 这一部分可以根据自己的需求修改
	// 如果不想调用 service 也可以直接在这里完成 RPC 的所有请求和响应
	resp, err := service.NewSayHelloService(ctx).Run(req)
	return resp, err
}
```
6. 在`app/frontend/biz/handler`和`app/frontend/biz/service`完成类似的代码。
7. 在`app/frontend/biz/router`中的模块中添加对应的路由。
8. 在`app/frontend/rpc/client.go`中添加rpc client的全局变量。具体写法可以参照其他变量修改。
9. 完善前端页面。

# Progress
## Framework
- [x] 前端路由（提供的路由接口框架基本完成，缺少前端页面的配合）
- [x] 后端各模块的rpc通信接口（提供的rpc通信接口已完成，位于handler.go文件）
- [x] 日志（日志初始化函数）
- [x] 数据库（数据库开启函数）
- [x] 加密算法
  - [x] 加盐的密码加密算法
- [x] config（基于etcd的参数保存、取用和watch）
- [x] 服务注册与发现（基于etcd的服务注册模块已经完成） 
  - [x] etcd
  - [ ] <del>consul</del>
- [x] 鉴权和认证
  - [x] jwt
  - [x] casbin
  - [ ] <del>satoken<del>
- [ ] 负载均衡
- [ ] 前端页面
## Auth
- [x] rpc通信
- [x] Token的生成与分发
- [x] Token验证，并通过casbin进行访问控制
- [ ] 服务注册
