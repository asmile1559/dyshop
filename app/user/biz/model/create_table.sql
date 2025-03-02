CREATE TABLE `users` (
    `user_id` bigint NOT NULL,
    `email` varchar(191) NOT NULL,
    `password` longtext,
    `phone` longtext,
    `name` longtext,
    `sign` longtext,
    `gender` longtext,
    `birthday` datetime(3) DEFAULT NULL,
    `role` longtext,
    `url` longtext,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `uni_users_email` (`email`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci