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
  `image_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
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


INSERT INTO monster (monster_id, name, description)
VALUES ('MON001', 'Slime', 'A blob of goo that can be surprisingly resilient.'),
       ('MON002', 'Goblin', 'A mischievous creature that loves to cause trouble.'),
       ('MON003', 'Dragon', 'A powerful and majestic creature that breathes fire.');

INSERT INTO item (item_id, name, image_url)
VALUES ('ITEM001', 'Potion', 'images/potion.png'),
       ('ITEM002', 'Sword', 'images/sword.jpg'),
       ('ITEM003', 'Armor', 'images/armor.gif');

INSERT INTO field (field_id, name, image_url)
VALUES ('FLD001', 'Forest', 'images/forest.jpg'),
       ('FLD002', 'Desert', 'images/desert.jpg'),
       ('FLD003', 'Mountain', 'images/mountain.png');

INSERT INTO product (product_id, name, publish_year, total_sales)
VALUES ('PRD001', 'Potion', '2023', '1000'),
       ('PRD002', 'Sword', '2020', '500'),
       ('PRD003', 'Armor', '2022', '200');

INSERT INTO part (part_id, monster_id, decription)
VALUES ('PRT001', 'MON001', 'Left arm'),
       ('PRT002', 'MON002', 'Sharp tooth'),
       ('PRT003', 'MON003', 'Fire breath');

INSERT INTO music (music_id, monster_id, name, image_url)
VALUES ('MSC001', 'MON001', 'Slime Symphony', 'images/music-slime.jpg'),
       ('MSC002', 'MON002', 'Goblin Groove', 'images/music-goblin.png'),
       ('MSC003', 'MON003', 'Dragons Ballad', 'images/music-dragon.gif');

INSERT INTO tribe (tribe_id, name_ja, name_en, description)
VALUES ('TRB001', 'ゴブリン族', 'Goblin Tribe', 'いたずら好きで集団で行動する'),
       ('TRB002', 'オーク族', 'Orc Tribe', '好戦的で力強い戦士集団'),
       ('TRB003', 'エルフ族', 'Elf Tribe', '長寿で優れた弓術を持つ');

INSERT INTO weakness (monster_id, aprt_id, fire, water, lightning, ice, dragon, slashing, blow, bullet)
VALUES ('MON001', 'PRT001', 'low', 'medium', 'high', 'low', 'immune', 'medium', 'high', 'low'),
       ('MON002', 'PRT002', 'medium', 'low', 'low', 'medium', 'low', 'high', 'low', 'medium'),
       ('MON003', 'PRT003', 'high', 'low', 'medium', 'high', 'high', 'low', 'medium', 'high');

INSERT INTO weapon (weapon_id, name, image_url, rarerity, attack, element_attack, shapness, critical, description)
VALUES ('WPN001', 'Wooden Sword', 'images/weapon_wooden_sword.jpg', 'Common', 'Low', 'None', 'Normal', 'Low', 'A basic sword made of wood.'),
       ('WPN002', 'Iron Sword', 'images/weapon_iron_sword.jpg', 'Uncommon', 'Medium', 'None', 'Normal', 'Medium', 'A sturdy sword made of iron.'),
       ('WPN003', 'Fire Sword', 'images/weapon_fire_sword.jpg', 'Rare', 'High', 'Fire', 'Normal', 'High', 'A sword imbued with fire that burns enemies.');
