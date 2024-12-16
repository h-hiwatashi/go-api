CREATE DATABASE  IF NOT EXISTS `go_clean_arch` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `go_clean_arch`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);

DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);