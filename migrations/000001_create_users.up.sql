create table users (
    id              bigserial       primary key,
    username        varchar(128)    not null unique,
    email           text            not null unique,
    full_name       text            not null,
    password_hash   text            not null
)

insert into users(username, email, full_name, password_hash)
values ('testusername111', 'testemail111@mail.com', 'John Doe', '$2a$10$eB3Axm6ikuREMCwYlxgrgOuEqjxL7r20ZIgaWziIL8JajzuXRQ6HW');