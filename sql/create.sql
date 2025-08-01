CREATE TABLE `comments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `content` text COLLATE utf8mb4_general_ci,
  `post_id` bigint unsigned DEFAULT NULL,
  `read_state` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



CREATE TABLE `links` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext COLLATE utf8mb4_general_ci,
  `url` longtext COLLATE utf8mb4_general_ci,
  `sort` bigint DEFAULT '0',
  `view` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_links_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `pages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` text COLLATE utf8mb4_general_ci,
  `body` longtext COLLATE utf8mb4_general_ci,
  `view` bigint DEFAULT NULL,
  `is_published` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `post_tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `post_id` bigint unsigned DEFAULT NULL,
  `tag_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_post_tag` (`post_id`,`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `posts` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` text COLLATE utf8mb4_general_ci,
  `body` longtext COLLATE utf8mb4_general_ci,
  `view` bigint DEFAULT NULL,
  `is_published` tinyint(1) DEFAULT NULL,
  `comment_total` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `smms_files` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `file_name` longtext COLLATE utf8mb4_general_ci,
  `store_name` longtext COLLATE utf8mb4_general_ci,
  `size` bigint DEFAULT NULL,
  `width` bigint DEFAULT NULL,
  `height` bigint DEFAULT NULL,
  `hash` longtext COLLATE utf8mb4_general_ci,
  `delete` longtext COLLATE utf8mb4_general_ci,
  `url` longtext COLLATE utf8mb4_general_ci,
  `path` longtext COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `subscribers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `verify_state` tinyint(1) DEFAULT '0',
  `subscribe_state` tinyint(1) DEFAULT '1',
  `out_time` datetime(3) DEFAULT NULL,
  `secret_key` longtext COLLATE utf8mb4_general_ci,
  `signature` longtext COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_subscribers_email` (`email`),
  KEY `idx_subscribers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` longtext COLLATE utf8mb4_general_ci,
  `total` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `email` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `telephone` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `verify_state` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '0',
  `secret_key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `out_time` datetime(3) DEFAULT NULL,
  `github_login_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `github_url` longtext COLLATE utf8mb4_general_ci,
  `is_admin` tinyint(1) DEFAULT NULL,
  `avatar_url` longtext COLLATE utf8mb4_general_ci,
  `nick_name` longtext COLLATE utf8mb4_general_ci,
  `lock_state` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`),
  UNIQUE KEY `idx_users_telephone` (`telephone`),
  UNIQUE KEY `idx_users_github_login_id` (`github_login_id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;