CREATE TABLE `activity` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day_midnight` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建日0点时间戳',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx-day_midnight` (`day_midnight`),
  KEY `idx-created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;