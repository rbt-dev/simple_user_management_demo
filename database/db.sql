CREATE DATABASE `simple_user_management_demo` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE `simple_user_management_demo`.`users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(128) NOT NULL,
  `firstname` varchar(128) DEFAULT NULL,
  `lastname` varchar(128) NOT NULL,
  `password` varchar(128) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `admin` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) DEFAULT CHARSET=utf8 COLLATE utf8_general_ci;

INSERT INTO `simple_user_management_demo`.`users` (`email`, `lastname`, `password`, `admin`) VALUES('admin@demo.com', 'admin', md5('admin'), 1);

