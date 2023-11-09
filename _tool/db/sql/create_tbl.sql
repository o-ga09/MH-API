DROP TABLE IF EXISTS `monster`;

CREATE TABLE `monster` (
    `id` INT(255) NOT NULL AUTO_INCREMENT,
    `monster_id` INT NOT NULL UNIQUE,
    `name` VARCHAR(255) NOT NULL,
    `desc` VARCHAR(255) NOT NULL,
    `location` VARCHAR(255),
    `category` VARCHAR(255),
    `title` VARCHAR(255),
    `weakness_attack` JSON,
    `weakness_element` JSON,
    PRIMARY KEY (`id`)
);

-- INSERT INTO `monster` (`name`,`desc`,`location`,`specify`,`weakness_attack`,`weakness_element`) VALUES ("ジンオウガ","霊峰/渓流に生息する電気を扱う牙竜種","霊峰","牙竜種","10 10 10 10 10","10 10 10 10 10");
-- INSERT INTO `monster` (`name`,`desc`,`location`,`specify`,`weakness_attack`,`weakness_element`) VALUES ("タマミツネ","渓流に生息する水を扱う海竜種","渓流","海竜種","10 10 10 10 10","10 10 10 10 10");

INSERT INTO monster (`monster_id`, `name`, `desc`, `location`, `category`, `title`, `weakness_attack`, `weakness_element`)
VALUES
    (1, 'Sample Monster 1', 'A sample monster with detailed elemental and attack weaknesses', 'Sample Location 2', 'Sample Category 2', 'Sample Title 2', '{"前脚": {"slashing": "10", "blow": "10", "bullet": "10"}, "尻尾": {"slashing": "10", "blow": "10", "bullet": "10"}, "後脚": {"slashing": "10", "blow": "10", "bullet": "10"}, "胴体": {"slashing": "10", "blow": "10", "bullet": "10"}, "頭部": {"slashing": "10", "blow": "10", "bullet": "10"}}', '{"前脚": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "尻尾": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "後脚": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "胴体": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "頭部": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}}'),
    (2, 'Sample Monster 2', 'A sample monster with detailed elemental and attack weaknesses', 'Sample Location 2', 'Sample Category 2', 'Sample Title 2', '{"前脚": {"slashing": "10", "blow": "10", "bullet": "10"}, "尻尾": {"slashing": "10", "blow": "10", "bullet": "10"}, "後脚": {"slashing": "10", "blow": "10", "bullet": "10"}, "胴体": {"slashing": "10", "blow": "10", "bullet": "10"}, "頭部": {"slashing": "10", "blow": "10", "bullet": "10"}}', '{"前脚": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "尻尾": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "後脚": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "胴体": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "頭部": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}}'),
    (3, 'Sample Monster 3', 'A sample monster with detailed elemental and attack weaknesses', 'Sample Location 2', 'Sample Category 2', 'Sample Title 2', '{"前脚": {"slashing": "10", "blow": "10", "bullet": "10"}, "尻尾": {"slashing": "10", "blow": "10", "bullet": "10"}, "後脚": {"slashing": "10", "blow": "10", "bullet": "10"}, "胴体": {"slashing": "10", "blow": "10", "bullet": "10"}, "頭部": {"slashing": "10", "blow": "10", "bullet": "10"}}', '{"前脚": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "尻尾": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "後脚": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "胴体": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}, "頭部": {"fire": "10", "water": "10", "lightning": "10", "ice": "10", "dragon": "10"}}');
