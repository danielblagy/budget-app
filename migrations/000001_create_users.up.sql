create table users (
    id              bigserial       primary key,
    username        varchar(128)    not null unique,
    email           text            not null unique,
    full_name       text            not null,
    password_hash   text            not null
);

insert into users(username, email, full_name, password_hash) values ('user1', 'user1@mymail.com', 'John', '');
insert into users(username, email, full_name, password_hash) values ('dblagy', 'dblagy@mymail.com', 'Daniel', '');
insert into users(username, email, full_name, password_hash) values ('randomname123', 'randomname123@mymail.com', 'Sally', '');
insert into users(username, email, full_name, password_hash) values ('user2', 'user2@mymail.com', 'Augustus', '');