-- +migrate Up
ALTER TABLE `item` ADD COLUMN `monster_id` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL;

-- +migrate Down
ALTER TABLE `item` DROP COLUMN `monster_id`;
