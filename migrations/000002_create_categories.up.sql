create type category_type as enum ('income', 'expense');

create table categories (
    id bigserial primary key,
    user_id varchar(128) not null,
    name varchar(128) not null,
    type category_type not null
);

create unique index categories_unique_idx on categories(user_id, type, name);

insert into categories(user_id, name, type) values ('dblagy', 'home', 'expense');
insert into categories(user_id, name, type) values ('dblagy', 'groceries', 'expense');
insert into categories(user_id, name, type) values ('dblagy', 'fun', 'expense');
insert into categories(user_id, name, type) values ('dblagy', 'eating out', 'expense');
insert into categories(user_id, name, type) values ('dblagy', 'day job', 'income');
insert into categories(user_id, name, type) values ('dblagy', 'vigilante', 'income');