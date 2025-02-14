CREATE TABLE `users` (
    `user_id` bigint NOT NULL,
    `email` varchar(191) NOT NULL,
    `password` longtext NOT NULL,
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `uni_users_email` (`email`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb3