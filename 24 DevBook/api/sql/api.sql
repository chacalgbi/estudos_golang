CREATE DATABASE IF NOT EXISTS `devbook`;
USE `devbook`;

DROP TABLE IF EXISTS `usuarios`;

CREATE TABLE `usuarios` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar(100) NOT NULL,
  `nick` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `senha` varchar(100) NOT NULL,
  `criadoEm` timestamp default current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE usuarios ALTER COLUMN senha TYPE VARCHAR(100);

ALTER TABLE usuarios ADD CONSTRAINT unique_nick UNIQUE(nick);
ALTER TABLE usuarios ADD CONSTRAINT unique_email UNIQUE(email);