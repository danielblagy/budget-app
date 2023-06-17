create type entry_type as enum ('income', 'expense');

create table entries(
    id bigserial primary key,
    user_id varchar(128) not null,
		category_id bigint not null,
		amount double precision not null,
		date timestamp not null,
		description text not null,
    type entry_type not null
);

create index entries_idx on entries(user_id, type);

insert into entries(user_id, category_id, amount, date, description, type)
values ('testusername111', '1', '120000.0', '2023-06-10', 'june salary', 'income');
insert into entries(user_id, category_id, amount, date, description, type)
values ('testusername111', '1', '5500.0', '2023-06-10', 'bonus', 'income');
insert into entries(user_id, category_id, amount, date, description, type)
values ('testusername111', '3', '453.19', '2023-06-08', 'BREAD', 'expense');
insert into entries(user_id, category_id, amount, date, description, type)
values ('testusername111', '3', '850.2', '2023-06-09', 'milk and eggs', 'expense');
insert into entries(user_id, category_id, amount, date, description, type)
values ('testusername111', '4', '600.0', '2023-06-09', 'cinema', 'expense');
