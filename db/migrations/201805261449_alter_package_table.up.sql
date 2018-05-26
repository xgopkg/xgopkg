ALTER TABLE `xgp_package` ADD `user_id` CHAR(36) NOT NULL AFTER `pkg_desc`;
CREATE TABLE IF NOT EXISTS `xgp_user`(
    `id` INT  UNSIGNED AUTO_INCREMENT,
    `user_id` CHAR(36) NOT NULL,
    `username` VARCHAR(100) NOT NULL,
    `password` VARCHAR(100) NOT NULL,
    `create_at` DATETIME NOT NULL,
    `update_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE uindex_user_id(user_id),
    UNIQUE uindex_username(username)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;