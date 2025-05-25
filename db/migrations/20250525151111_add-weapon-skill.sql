-- +migrate Up
ALTER TABLE `weapon` ADD COLUMN `skill_1` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL;
ALTER TABLE `weapon` ADD COLUMN `skill_2` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL;
-- +migrate Down
ALTER TABLE `weapon` DROP COLUMN `skill_1`;
ALTER TABLE `weapon` DROP COLUMN `skill_2`;
