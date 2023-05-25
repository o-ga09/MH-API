CREATE TABLE IF NOT EXISTS `monsters` (
    `id` INT(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `desc` TEXT NOT NULL,
    `location` VARCHAR(255),
    `specify` VARCHAR(255),
    `weakness_attack` VARCHAR(255),
    `weakness_element` VARCHAR(255),
    PRIMARY KEY (`id`)
);