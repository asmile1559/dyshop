CREATE TABLE `users` (
    `user_id` bigint NOT NULL,
    `email` varchar(191) NOT NULL,
    `password` longtext NOT NULL,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `uni_users_email` (`email`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3

CREATE TABLE `users` (
    `user_id` BIGINT NOT NULL PRIMARY KEY,    -- 用户ID（雪花算法生成）
    `email` VARCHAR(191) NOT NULL UNIQUE,     -- 邮箱（唯一约束）
    `password` VARCHAR(255) NOT NULL,         -- 加密存储的密码
    `phone` VARCHAR(20) DEFAULT NULL,         -- 手机号（可选）
    `name` VARCHAR(100) DEFAULT NULL,         -- 用户昵称
    `sign` VARCHAR(255) DEFAULT NULL,         -- 个性签名
    `gender` ENUM('男', '女', '不愿意公开') DEFAULT '不愿意公开', -- 性别（男、女、不愿意公开）
    `birthday` DATE DEFAULT NULL,             -- 生日
    `role` ENUM('user', 'merchant') DEFAULT 'user', -- 角色（普通用户、商户）
    `url` VARCHAR(255) DEFAULT NULL   -- 头像存储URL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
