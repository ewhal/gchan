CREATE TABLE IF NOT EXISTS `boards` (
  `uri` varchar(58) CHARACTER SET utf8 NOT NULL,
  `title` tinytext NOT NULL,
  `subtitle` tinytext,
  -- `indexed` boolean default true,
  PRIMARY KEY (`uri`)
);
INSERT INTO `boards` VALUES
('b', 'Random', NULL);

