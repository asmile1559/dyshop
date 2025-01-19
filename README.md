# dyshop

## project architecture

- /
    - app/: microservice application
        - app1/:
            - biz(business)/: 业务代码写在这里
                - dal(data access layer)/: 数据库访问
                - service/: 具体要实现的逻辑功能, service会调用handler的代码
                - model/: 成员结构
            - conf/: 存放配置文件
            - middleware/: 中间件函数，存放逻辑中间件
            - script/: 脚本
            - utils/: 存放工具函数
            - docker-compose.yaml: docker容器启动
            - handler.go: grpc实现接口
            - main.go: app入口
            - .env: 环境变量
            - go.mod
        - app2/:
            - ...
    - proto/: 存放proto文件
      - frontend: 前端的proto文件
      - backend: 后端的proto文件
    - pb:根据proto文件生成的grpc文件
      - frontend: 前端对应的grpc文件
      - backend: 后端对应的grpc文件
    - Makefile: 常用的构建命令
    - docker-compose.yaml: docker容器启动
    - go.work: work目录
    - .gitignore
    
## Communication:

1. 外部请求通过 RESTful API 发送至 gateway(frontend)
2. 当一个app启动时，会向注册中心进行注册
3. 外部对某一个服务的访问，经过 gateway 进行转发
4. 内部服务之间的访问，通过注册中心找到对应的服务？

## 框架需要搭建什么？
1. 前端的访问接口？
    - [ ] auth
    - [ ] cart
    - [ ] checkout
    - [ ] order
    - [ ] payment
    - [ ] product
    - [ ] user
2. 后端各模块的rpc通信接口？