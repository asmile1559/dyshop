# dyshop

## user注册和登录

```bash
# 1. 下载所需的依赖
go mod tidy

# 2. 运行 mysql数据库
cd docker-compose/user-compose && docker compose up -d

# 3. 在mysql数据库中创建对应的数据库和表
docker exec -it user_compose-mysql-1 bash

mysql -u root -p root
create database user;
use user;
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `avatar_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `role` enum('user','admin') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'user',
  `status` tinyint NULL DEFAULT 1 COMMENT '1: active, 0: inactive, -1: deleted',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username` ASC) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE,
  INDEX `idx_email`(`email` ASC) USING BTREE,
  INDEX `idx_phone`(`phone` ASC) USING BTREE,
  INDEX `idx_status_deleted`(`status` ASC, `deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

# 4. 在另一个终端中运行
go run cmd/server/main.go

# 5. 在浏览器中
# 访问注册
# http://192.168.191.130:10166/user/register
#         ↑           ↑    ↑    ↑     ↑
#         这一段是ip地址  默认端口 路由  注册
# 注册成功可以在数据库中查到对应的用户信息
# 访问登录
# http://192.168.191.130:10166/user/login
#         ↑           ↑    ↑    ↑     ↑
#         这一段是ip地址  默认端口 路由   登录
# 返回 {"message":{"token":"xxx","userId":xxx}} 时，说明登录成功，没有错误
# 返回 {"err": "xxx"} 时说明登录失败
```

