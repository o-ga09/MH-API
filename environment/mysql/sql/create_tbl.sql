DROP TABLE IF EXISTS `field`;
DROP TABLE IF EXISTS `item`;
DROP TABLE IF EXISTS `monster`;
DROP TABLE IF EXISTS `music`;
DROP TABLE IF EXISTS `part`;
DROP TABLE IF EXISTS `product`;
DROP TABLE IF EXISTS `tribe`;
DROP TABLE IF EXISTS `weakness`;
DROP TABLE IF EXISTS `weapon`;

CREATE TABLE `monster` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`monster_id`),
  KEY `idx_monster_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `item_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `imege_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`item_id`),
  KEY `idx_item_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `field` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `field_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `image_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`field_id`),
  KEY `idx_field_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `product_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `publish_year` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `total_sales` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`product_id`),
  KEY `idx_product_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `part` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `part_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `decription` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`part_id`),
  KEY `idx_part_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `music` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `music_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `image_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`music_id`),
  KEY `idx_music_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `tribe` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `tribe_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name_ja` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name_en` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`tribe_id`),
  KEY `idx_tribe_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `weakness` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `aprt_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `fire` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `water` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lightning` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ice` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `dragon` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `slashing` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `blow` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bullet` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `first_weak_attack` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `second_weak_attack` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `first_weak_element` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `second_weak_element` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_weakness_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `weapon` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `weapon_id` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `image_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `rarerity` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `attack` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `element_attack` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `shapness` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `critical` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`weapon_id`),
  KEY `idx_weapon_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
