
SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `msgs`
-- ----------------------------
-- DROP TABLE IF EXISTS `msgs`;

CREATE TABLE `msgs` (
	`topic` VARCHAR (255) NOT NULL,
	`id` VARCHAR (255) NOT NULL,
	`original_id` VARCHAR (255) NOT NULL,
	`type` INT (11) NOT NULL,
	`payload` BLOB,
	`insert_time` BIGINT (22) NOT NULL,
	PRIMARY KEY (`topic`, `id`, `original_id`),
	KEY `id_index` (`id`) USING BTREE,
	KEY `original_id_index` (`original_id`) USING BTREE,
	KEY `insert_time_index` (`insert_time`) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;

-- SET FOREIGN_KEY_CHECKS = 1;
