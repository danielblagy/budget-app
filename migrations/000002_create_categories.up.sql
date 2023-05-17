create table categories (
    id bigserial primary key,
    user_id varchar(128) not null,
    name varchar(128) not null
);

create unique index categories_unique_idx on categories(user_id, name);

insert into categories(user_id, name) values ('dblagy', 'home');
insert into categories(user_id, name) values ('dblagy', 'groceries');
insert into categories(user_id, name) values ('dblagy', 'fun');
insert into categories(user_id, name) values ('dblagy', 'eating out');