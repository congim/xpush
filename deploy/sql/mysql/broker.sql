
SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `msgs`
-- ----------------------------
-- DROP TABLE IF EXISTS `msgs`;

CREATE TABLE `msgs` (
  `topic` varchar(255) NOT NULL,
  `id` varchar(255) NOT NULL,
  `original_id` varchar(255) NOT NULL,
  `type` int(11) NOT NULL,
  `payload` blob,
  `insert_time` datetime NOT NULL,
  PRIMARY KEY (`topic`,`id`,`original_id`),
  KEY `id_index` (`id`) USING BTREE,
  KEY `original_id_index` (`original_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- SET FOREIGN_KEY_CHECKS = 1;
