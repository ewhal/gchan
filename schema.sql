CREATE TABLE `boards` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `board` longtext,
  `title` longtext,
  `subtitle` longtext,
  `description` longtext,
  PRIMARY KEY (`id`)
);

CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `board` longtext,
  `postnum` int(20),
  `op` boolean,
  `title` longtext,
  `name` longtext,
  `email` longtext,
  `usermode` longtext,
  `post` longtext,
  `files` longtext,
  `created` TIMESTAMP,
  `thread` longtext,
  PRIMARY KEY (`id`)
);

CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `email` VARCHAR(320),
  `password` CHAR(76),
  `boards` CHAR(76),
  `level` CHAR(76),
  PRIMARY KEY (`id`)
);

