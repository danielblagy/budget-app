create table users (
    id      bigserial   primary key,
    handle  text        not null
);

insert into users(handle) values ('user1');
insert into users(handle) values ('dblagy');
insert into users(handle) values ('doe');