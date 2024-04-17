insert into usuarios (nome, nick, email, senha) values
('usuario1', 'nick1', 'email1@gmail.com', '$2a$10$TK.dZkyh9NHTINhDJsk9B.CO6WwVs1b4l1XFHxjcC7k3YIQ5L0JeK'),
('usuario2', 'nick2', 'email2@gmail.com', '$2a$10$TK.dZkyh9NHTINhDJsk9B.CO6WwVs1b4l1XFHxjcC7k3YIQ5L0JeK'),
('usuario3', 'nick3', 'email3@gmail.com', '$2a$10$TK.dZkyh9NHTINhDJsk9B.CO6WwVs1b4l1XFHxjcC7k3YIQ5L0JeK'),
('usuario4', 'nick4', 'email4@gmail.com', '$2a$10$TK.dZkyh9NHTINhDJsk9B.CO6WwVs1b4l1XFHxjcC7k3YIQ5L0JeK');

insert into seguidores (usuario_id, seguidor_id) values
(1, 2),
(1, 3),
(2, 1),
(2, 3),
(3, 1),
(3, 2),
(4, 1),
(4, 2),
(4, 3);