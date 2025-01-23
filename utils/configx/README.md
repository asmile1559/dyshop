# etcd 使用 Docker Compose 的设置

## 概述
如何通过 Docker Compose 部署单节点 etcd，并直接在容器内与 etcd 进行交互来管理配置。

## Docker Compose 配置
以下是 `docker-compose.yml` 文件：

```yaml
services:
  etcd:
    image: quay.io/coreos/etcd:latest
    container_name: etcd_server
    ports:
      - "2379:2379"
      - "2380:2380"
    command: |
      /usr/local/bin/etcd \
      --name etcd0 \
      --data-dir /etcd-data \
      --advertise-client-urls http://0.0.0.0:2379 \
      --listen-client-urls http://0.0.0.0:2379
```

### 关键点
- **端口映射：**
  - etcd 的客户端访问端口映射到 `2379`。
  - etcd 的集群通信端口映射到 `2380`。
- **命令参数：**
  - `--advertise-client-urls` 和 `--listen-client-urls` 设置为允许容器内外访问。

## 部署步骤

1. 将上述 `docker-compose.yml` 保存到工作目录。

2. 执行以下命令启动服务：

   ```bash
   docker-compose up -d
   ```

3. 验证容器是否运行：

   ```bash
   docker ps
   ```

## 与 etcd 交互 (以hello-service为例)

1. 在 `deploy/` 目录下使用 `docker-compose exec`直接执行命令：

   ```bash
   # 写入键值对
   docker compose exec -e ETCDCTL_API=3 etcd etcdctl put /config/hello-service/name "Alice"

   # 获取键值对
   docker compose exec -e ETCDCTL_API=3 etcd etcdctl get /config/hello-service --prefix
   ```
2. 使用 `utils/configx` 里的相关函数:
    - `utils/configx/config_store.go`: 存储配置到 etcd
    - `utils/configx/config_fetch.go`: 从 etcd 获取配置

## 验证

1. 使用 `get` 命令确保键值对已存储到 etcd。
2. 将应用程序集成到 etcd 中，读取并监听配置键（如 `/config/hello-service/name`）。
3. 测试动态更新：
   - 使用 `etcdctl` 修改键值。
   - 观察应用程序日志，确认其检测到配置更改并相应更新（参考utils/example/registry_client/main.go）。


