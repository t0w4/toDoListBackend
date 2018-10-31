# mysql -u root < init.sql
DROP DATABASE IF EXISTS `todoList`;
CREATE DATABASE `todoList` CHARACTER SET utf8mb4;

CREATE TABLE `todoList`.`tasks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL DEFAULT '',
  `title` varchar(255) NOT NULL DEFAULT '',
  `detail` varchar(255) DEFAULT '' COMMENT 'titleだけでもOK',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
