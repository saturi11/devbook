--senha = 1234567

insert into usuarios (nome ,nick, email, senha)
values
("usuario 1","usuario_1","usuario1@gmail.com","$2a$10$4arbcmkez8beiYo2qgLwDuMTvQm0ICbgCsEH4s4OawYEITiZUN1hS"), 
("usuario 2","usuario_2","usuario2@gmail.com","$2a$10$4arbcmkez8beiYo2qgLwDuMTvQm0ICbgCsEH4s4OawYEITiZUN1hS"),
("usuario 3","usuario_3","usuario3@gmail.com","$2a$10$4arbcmkez8beiYo2qgLwDuMTvQm0ICbgCsEH4s4OawYEITiZUN1hS"),
("usuario 4","usuario_4","usuario4@gmail.com","$2a$10$4arbcmkez8beiYo2qgLwDuMTvQm0ICbgCsEH4s4OawYEITiZUN1hS");


insert into seguidores (usuario_id ,seguidor_id)
values
(1, 2), 
(3, 1),
(4, 2),
(2, 1);


insert into publicacoes(titulo, conteudo, autor_id)
values
("publicacao do user 1","essa e a publi do user 1",1),
("publicacao do user 2","essa e a publi do user 2",2),
("publicacao do user 3","essa e a publi do user 3",3),
("publicacao do user 4","essa e a publi do user 4",4);