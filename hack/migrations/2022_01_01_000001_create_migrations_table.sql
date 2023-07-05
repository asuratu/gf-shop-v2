CREATE TABLE `migrations` (
     `id` int unsigned NOT NULL AUTO_INCREMENT,
     `migration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
     `created_at` timestamp NULL DEFAULT NULL,
     PRIMARY KEY (`id`) USING BTREE,
     UNIQUE KEY `migrations_migration_unique` (`migration`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=534 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='迁移记录表';

INSERT INTO `gf-shop`.`migrations` (`id`, `migration`, `created_at`) VALUES (1, '2022_01_01_000001_create_migrations_table', '2023-07-05 09:43:27');
