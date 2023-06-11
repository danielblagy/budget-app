create type category_type as enum ('income', 'expense');

create table categories (
    id bigserial primary key,
    user_id varchar(128) not null,
    name varchar(128) not null,
    type category_type not null
);

create unique index categories_unique_idx on categories(user_id, type, name);

insert into categories(user_id, name, type) values ('testusername111', 'day job', 'income');
insert into categories(user_id, name, type) values ('testusername111', 'vigilante', 'income');
insert into categories(user_id, name, type) values ('testusername111', 'groceries', 'expense');
insert into categories(user_id, name, type) values ('testusername111', 'fun', 'expense');
insert into categories(user_id, name, type) values ('testusername111', 'eating out', 'expense');