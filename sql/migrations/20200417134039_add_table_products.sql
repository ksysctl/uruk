CREATE TABLE IF NOT EXISTS `products` (
  `id` INT(10) UNSIGNED AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL DEFAULT '',
  `code` VARCHAR(255) NOT NULL DEFAULT '',
  `price` DOUBLE NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `code_unique` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=1;
