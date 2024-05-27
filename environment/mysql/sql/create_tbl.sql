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
  `monster_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`,`monster_id`),
  KEY `idx_monster_deleted_at` (`deleted_at`),
  KEY `idx_monster_monster_id` (`monster_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `item` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `item_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description`  varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `image_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`,`item_id`),
  KEY `idx_item_deleted_at` (`deleted_at`),
  KEY `fk_monster_item` (`monster_id`),
  CONSTRAINT `fk_monster_item` FOREIGN KEY (`monster_id`) REFERENCES `monster` (`monster_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `field` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `field_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`,`field_id`),
  KEY `idx_field_deleted_at` (`deleted_at`),
  KEY `fk_monster_field` (`monster_id`),
  CONSTRAINT `fk_monster_field` FOREIGN KEY (`monster_id`) REFERENCES `monster` (`monster_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `product_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `publish_year` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `total_sales` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`,`product_id`),
  KEY `idx_product_deleted_at` (`deleted_at`),
  KEY `fk_monster_product` (`monster_id`),
  CONSTRAINT `fk_monster_product` FOREIGN KEY (`monster_id`) REFERENCES `monster` (`monster_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `part` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `part_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
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
  `tribe_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name_ja` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name_en` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`,`tribe_id`),
  KEY `idx_tribe_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_monster_tribe` FOREIGN KEY (`monster_id`) REFERENCES `monster` (`monster_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `weakness` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `part_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `fire` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `water` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `lightning` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ice` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `dragon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `slashing` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `blow` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `bullet` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `first_weak_attack` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `second_weak_attack` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `first_weak_element` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `second_weak_element` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_weakness_deleted_at` (`deleted_at`),
  KEY `fk_monster_weakness` (`monster_id`),
  CONSTRAINT `fk_monster_weakness` FOREIGN KEY (`monster_id`) REFERENCES `monster` (`monster_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

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

CREATE TABLE `ranking` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `monster_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ranking` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `vote_year` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_monster_deleted_at` (`deleted_at`),
  KEY `idx_monster_monster_id` (`monster_id`),
  CONSTRAINT `fk_monster_ranking` FOREIGN KEY (`monster_id`) REFERENCES `monster` (`monster_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO monster (monster_id, name, description)
VALUES ('0000000001', 'リオレウス', '空の王者'),
       ('0000000002', 'リオレイア', '陸の女王'),
       ('0000000003', 'ティガレックス', '絶対強者');

INSERT INTO item (item_id,monster_id, name, image_url)
VALUES ('ITM001','0000000001', '火竜の翼爪', 'images/image1.jpg'),
       ('ITM002','0000000001', '火竜の鱗', 'images/image2.jpg'),
       ('ITM003','0000000002', '雌火竜の紅玉', 'images/image3.jpg'),
       ('ITM004','0000000002', '雌火竜の甲殻', 'images/image4.jpg'),
       ('ITM005','0000000003', '轟竜の甲殻', 'images/image5.jpg');

INSERT INTO field (field_id,monster_id, name, image_url)
VALUES ('FLD001','0000000001', '森丘', 'images/forest.jpg'),
       ('FLD002','0000000001', '渓流', 'images/desert.jpg'),
       ('FLD002','0000000001', '大社跡', 'images/desert.jpg'),
       ('FLD002','0000000002', '溶岩洞', 'images/desert.jpg'),
       ('FLD003','0000000003', '砂原', 'images/mountain.png');

INSERT INTO product (product_id,monster_id, name, publish_year, total_sales)
VALUES ('PRD001', '0000000001','MH', '2004', '1000'),
       ('PRD002', '0000000001','MH2', '2001', '500'),
       ('PRD003', '0000000001','MH3', '2009', '500'),
       ('PRD004', '0000000002','MHR', '2021', '500'),
       ('PRD005', '0000000003','MHW', '2018', '200');

INSERT INTO part (part_id, name, monster_id, decription)
VALUES ('PRT001','頭部', '0000000001', 'Left arm'),
       ('PRT002','腹部', '0000000001', 'Left arm'),
       ('PRT003','前脚', '0000000001', 'Left arm'),
       ('PRT004','後脚', '0000000001', 'Left arm'),
       ('PRT005','尻尾', '0000000001', 'Left arm'),
       ('PRT001','頭部', '0000000002', 'Sharp tooth'),
       ('PRT002','腹部', '0000000002', 'Sharp tooth'),
       ('PRT003','前脚', '0000000002', 'Sharp tooth'),
       ('PRT004','後脚', '0000000002', 'Sharp tooth'),
       ('PRT005','尻尾', '0000000002', 'Sharp tooth'),
       ('PRT001','頭部', '0000000003', 'Fire breath'),
       ('PRT002','腹部', '0000000003', 'Fire breath'),
       ('PRT003','前脚', '0000000003', 'Fire breath'),
       ('PRT004','後脚', '0000000003', 'Fire breath'),
       ('PRT005','尻尾', '0000000003', 'Fire breath');

INSERT INTO music (music_id, monster_id, name, image_url)
VALUES ('MSC001', '0000000001', 'Slime Symphony', 'images/music-slime.jpg'),
       ('MSC002', '0000000002', 'Goblin Groove', 'images/music-goblin.png'),
       ('MSC003', '0000000003', 'Dragons Ballad', 'images/music-dragon.gif');

INSERT INTO tribe (tribe_id,monster_id, name_ja, name_en, description)
VALUES ('TRB001','0000000001', '飛竜種', 'Wybarn', '空を飛ぶ竜種'),
       ('TRB002','0000000002', '飛竜種', 'Wybarn', '空を飛ぶ竜種'),
       ('TRB003','0000000003', '飛竜種', 'Wybarn', '空を飛ぶ竜種');

INSERT INTO weakness (monster_id, part_id, fire, water, lightning, ice, dragon, slashing, blow, bullet,first_weak_attack, second_weak_attack, first_weak_element, second_weak_element)
VALUES ('0000000001', 'PRT001', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000001', 'PRT002', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000001', 'PRT003', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000001', 'PRT004', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000001', 'PRT005', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000002', 'PRT001', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000002', 'PRT002', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000002', 'PRT003', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000002', 'PRT004', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000002', 'PRT005', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000003', 'PRT001', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000003', 'PRT002', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000003', 'PRT003', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000003', 'PRT004', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水'),
       ('0000000003', 'PRT005', '45', '45', '45', '45', '45', '45', '45', '45','斬','弾','火','水');

INSERT INTO weapon (weapon_id, name, image_url, rarerity, attack, element_attack, shapness, critical, description)
VALUES ('WPN001', 'Wooden Sword', 'images/weapon_wooden_sword.jpg', 'Common', 'Low', 'None', 'Normal', 'Low', 'A basic sword made of wood.'),
       ('WPN002', 'Iron Sword', 'images/weapon_iron_sword.jpg', 'Uncommon', 'Medium', 'None', 'Normal', 'Medium', 'A sturdy sword made of iron.'),
       ('WPN003', 'Fire Sword', 'images/weapon_fire_sword.jpg', 'Rare', 'High', 'Fire', 'Normal', 'High', 'A sword imbued with fire that burns enemies.');

INSERT INTO ranking (monster_id, ranking, vote_year)
VALUES ('0000000001', '1', '2024/03/12'),
       ('0000000002', '2', '2024/03/12'),
       ('0000000003', '3', '2024/03/12');
