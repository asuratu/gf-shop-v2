-- Date: 2023-07-05 09:43:27
-- 创建手工位表 - position_info

-- 迁移记录
INSERT INTO `gf-shop`.`migrations` (`migration`, `created_at`) VALUES ('2023_07_05_093049_create_position_info_table', '2023-07-05 09:43:27');

-- 表结构 - position_info
CREATE TABLE `position_info` (
     `id` int NOT NULL AUTO_INCREMENT,
     `pic_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图片链接',
     `goods_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
     `link` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '跳转链接',
     `sort` tinyint NOT NULL DEFAULT '0' COMMENT '排序',
     `goods_id` int NOT NULL DEFAULT '0' COMMENT '商品id',
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='手工位表';

-- 插入数据
INSERT INTO `gf-shop`.`position_info` (`id`, `pic_url`, `goods_name`, `link`, `sort`, `goods_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'https://images.zsxq.com/FgdL08hVmh-40_e12vh-ifbXpGxB?e=2000966400', '测试', 'https://articles.zsxq.com/id_wd15wsegvow1.html', 0, 1, '2022-11-18 17:44:07', '2022-11-18 17:44:07', '2022-11-18 17:44:59');
INSERT INTO `gf-shop`.`position_info` (`id`, `pic_url`, `goods_name`, `link`, `sort`, `goods_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'http://dummyimage.com/120x60', '它', 'http://yypbiyi.nz/hbcygud', 1, 81, '2023-03-20 11:15:04', '2023-03-20 11:15:04', '2023-03-20 15:58:49');
INSERT INTO `gf-shop`.`position_info` (`id`, `pic_url`, `goods_name`, `link`, `sort`, `goods_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'http://dummyimage.com/160x600', '住', 'http://ollq.hn/kcsnsqoh', 1, 68, '2023-03-20 11:15:05', '2023-03-20 17:25:01', NULL);
INSERT INTO `gf-shop`.`position_info` (`id`, `pic_url`, `goods_name`, `link`, `sort`, `goods_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'http://dummyimage.com/160x600', '万', 'http://jynbisii.cu/munsfcey', 1, 93, '2023-03-20 11:15:27', '2023-03-20 11:15:27', NULL);
INSERT INTO `gf-shop`.`position_info` (`id`, `pic_url`, `goods_name`, `link`, `sort`, `goods_id`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'http://dummyimage.com/336x280', '低', 'http://slcnbfezx.dm/gkoomhxgj', 1, 96, '2023-03-20 11:15:30', '2023-03-20 11:15:30', NULL);
