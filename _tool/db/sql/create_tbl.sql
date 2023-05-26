DROP TABLE IF EXISTS `monster`;

CREATE TABLE `monster` (
    `id` INT(255) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `desc` VARCHAR(255) NOT NULL,
    `location` VARCHAR(255),
    `specify` VARCHAR(255),
    `weakness_attack` VARCHAR(255),
    `weakness_element` VARCHAR(255),
    PRIMARY KEY (`id`)
);

INSERT INTO `monster` (`name`,`desc`,`location`,`specify`,`weakness_attack`,`weakness_element`) VALUES ("ジンオウガ","霊峰/渓流に生息する電気を扱う牙竜種","霊峰","牙竜種","10 10 10 10 10","10 10 10 10 10");
INSERT INTO `monster` (`name`,`desc`,`location`,`specify`,`weakness_attack`,`weakness_element`) VALUES ("タマミツネ","渓流に生息する水を扱う海竜種","渓流","海竜種","10 10 10 10 10","10 10 10 10 10");