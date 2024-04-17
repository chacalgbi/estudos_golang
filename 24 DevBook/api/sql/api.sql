# mysql -u lucas -p Entra no mysql via terminal
# use devbook; Seleciona o banco de dados
# desc usuarios; Mostra a estrutura da tabela usuarios

CREATE DATABASE IF NOT EXISTS `devbook`;
USE `devbook`;

DROP TABLE IF EXISTS `usuarios`;
DROP TABLE IF EXISTS `seguidores`;

CREATE TABLE `usuarios` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nome` varchar(100) NOT NULL,
  `nick` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `senha` varchar(100) NOT NULL,
  `criadoEm` timestamp default current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `seguidores`(
  `usuario_id` int(11) NOT NULL, FOREIGN KEY (`usuario_id`) REFERENCES usuarios(`id`) ON DELETE CASCADE,
  `seguidor_id` int(11) NOT NULL, FOREIGN KEY (`seguidor_id`) REFERENCES usuarios(`id`) ON DELETE CASCADE,

  PRIMARY KEY (`usuario_id`, `seguidor_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE usuarios ALTER COLUMN senha TYPE VARCHAR(100);

ALTER TABLE usuarios ADD CONSTRAINT unique_nick UNIQUE(nick);
ALTER TABLE usuarios ADD CONSTRAINT unique_email UNIQUE(email);