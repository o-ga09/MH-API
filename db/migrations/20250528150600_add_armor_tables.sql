-- +migrate Up
CREATE TABLE IF NOT EXISTS `armor` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `armor_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slot` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `defense` int DEFAULT NULL,
  `fire_resistance` int DEFAULT NULL,
  `water_resistance` int DEFAULT NULL,
  `lightning_resistance` int DEFAULT NULL,
  `ice_resistance` int DEFAULT NULL,
  `dragon_resistance` int DEFAULT NULL,
  PRIMARY KEY (`id`,`armor_id`),
  KEY `idx_armor_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `armor_skill` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `armor_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `skill_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `skill_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_armor_skill_deleted_at` (`deleted_at`),
  KEY `fk_armor_skill` (`armor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `armor_required_item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `armor_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `item_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_armor_required_item_deleted_at` (`deleted_at`),
  KEY `fk_armor_required_item` (`armor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS `armor_required_item`;
DROP TABLE IF EXISTS `armor_skill`;
DROP TABLE IF EXISTS `armor`;
