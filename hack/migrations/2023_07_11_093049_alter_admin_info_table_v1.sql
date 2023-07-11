-- Date: 2023-07-11 14:25:53
-- 更新表结构, 增加 deleted_at 字段 - admin_info

-- 迁移记录
INSERT INTO `gf-shop`.`migrations` (`migration`, `created_at`) VALUES ('2023_07_11_093049_2023_07_11_093049_alter_admin_info_table_v1', '2023-07-11 14:25:53');

-- 表结构 - admin_info
ALTER TABLE `admin_info` ADD COLUMN `deleted_at` datetime DEFAULT NULL COMMENT '删除时间' AFTER `updated_at`;
