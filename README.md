# dyshop

## 目录
- [Workflow](#workflow)
- [How to Run](#how-to-run)
- [Project Architecture](#project-architecture)
- [Communication](#communication)
- [如何利用gin和grpc实现一个服务](#如何利用gin和grpc实现一个服务)
- [进度](#进度)

# Workflow
1. 将代码克隆到本地。
```bash
git clone git@github.com:asmile1559/dyshop.git
````
2. 新建一个分支
```bash
# example需要替换为自己负责的模块
git branch feat/example
git checkout feat/example
```

3. 安装 `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`.
```bash
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

4. 生成proto文件（可以不生成）
```bash
make gen-frontend-proto
make gen-backend-proto
```

5. 开始编写代码吧。

当开始写某个微服务的代码时，例如`auth`，`app/auth/service/`中已经提供两个待实现的接口。也可以在`app/auth/handler.go`文件中编写grpc代码。
尽量在service中编写，方便维护。

**检测功能有三种方式。** \
①直接在service的test中测试函数中编写测试样例，验证是否可以达到所需结果。\
②在对应模块的 cmd/clinet/main.go 中修改调用函数和参数，验证是否可以达到所需结果。\
③启动frontend，使用postman或者浏览器访问路由，验证是否可以达到所需结果。

6. 推送分支
```bash
git add .
git commit -m "example commit"
git push origin feat/example 
```
7. 发起一个PR请求

# How to Run
```bash
# 以 user 微服务为例
# 如果采用方式 ① 
#   直接运行 app/user/biz/service/xxx_test.go 即可（需要自己完成调用）
# 如果采用方式 ② 
# 1. 先启动微服务
cd app/user && go run . 
# 2. 启动一个新的终端，执行本地的 rpc 调用（需要自己决定如何调用）
cd app/user/cmd/client && go run .
# 如果采用方式 ③
# 1. 先启动微服务
cd app/user && go run . 
# 2. 启动一个新的终端
cd app/frontend && go run .
# 3. 使用 postman 或者 浏览器 或者 curl 进行访问（需要自己完成前端页面和路由逻辑）
curl -X POST \
          -H "Content-Type: application/json" \
          -d '{"email":"123@abc.com", "password": "123456", "confirm_password": "123456"}' \
          localhost:12166 
```

# Project Architecture

- /
    - app/: microservice application
        - app1/:
            - biz(business)/: 业务代码写在这里
                - dal(data access layer)/: 数据库访问
                - service/: 具体要实现的逻辑功能, service会调用handler的代码
                - model/: 成员结构
            - cmd: 存放一个grpc client的测试程序，用于测试grpc是否可用
            - conf/: 存放配置文件
            - middleware/: 中间件函数，存放逻辑中间件
            - script/: 脚本
            - utils/: 存放工具函数
            - .env: 环境变量
            - docker-compose.yaml: docker容器启动
            - go.mod
            - handler.go: grpc实现接口
            - main.go: app入口
        - app2/:
            - ...
    - deploy: 部署使用的文件
    - pb:根据proto文件生成的grpc文件
      - frontend: 前端对应的grpc文件
      - backend: 后端对应的grpc文件
    - proto/: 存放proto文件
      - frontend: 前端的proto文件
      - backend: 后端的proto文件
    - utils: 存放项目公用工具函数
      - configx: 配置管理相关函数
      - db: 数据库函数
      - example: 工具函数的实例
      - filex: file相关函数
      - jwt: 鉴权相关函数
      - logx: 日志相关函数
      - registryx: 注册相关函数
    - .gitignore
    - dyshop.postman_collection.json: postman的请求文件，参考使用
    - go.work: work目录
    - Makefile: 常用的构建命令
    - README.md: 说明文档
    
# Communication:

1. 外部请求通过 RESTful API 发送至 gateway(frontend)
2. 当一个app启动时，会向注册中心进行注册
3. 外部对某一个服务的访问，经过 gateway 进行转发
4. 内部服务之间的访问，通过注册中心找到对应的服务？

# 如何利用gin和grpc实现一个服务
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

# 进度
- [x] 前端路由（提供的路由接口框架基本完成，缺少前端页面的配合）
- [x] 后端各模块的rpc通信接口（提供的rpc通信接口已完成，位于handler.go文件）
- [x] 服务注册与发现（基于etcd的服务注册模块已经完成） 
  - [ ] <del>consul</del>
  - [x] etcd
- [x] 鉴权和认证
  - [x] jwt
  - [x] casbin
  - [ ] <del>satoken<del>
- [ ] 前端页面
  - [ ] html
  - [ ] js
  - [ ] template
