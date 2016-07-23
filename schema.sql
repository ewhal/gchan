CREATE TABLE `board` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `title` longtext,
  `subtitle` longtext,
  `description` longtext,
  PRIMARY KEY (`id`)
);

CREATE TABLE `threads` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `board` longtext,
  `title` longtext,
  `name` longtext,
  `email` longtext,
  `usermode` longtext,
  `post` longtext,
  `files` longtext,
  `created` DATESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `posts` (
  `id` int(10) unsigned NOT NULL auto_increment,
  `board` longtext,
  `title` longtext,
  `name` longtext,
  `email` longtext,
  `usermode` longtext,
  `post` longtext,
  `files` longtext,
  `created` DATESTAMP,
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

