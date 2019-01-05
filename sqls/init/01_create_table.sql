CREATE TABLE `todoList`.`tasks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL DEFAULT '',
  `title` varchar(255) NOT NULL DEFAULT '',
  `detail` varchar(255) DEFAULT '' COMMENT 'titleだけでもOK',
  `status` varchar(255) NOT NULL DEFAULT 'todo' COMMENT 'todo, in_progress. in_review, done',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
