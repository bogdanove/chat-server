-- +goose Up
create table chat_log
(
    id      serial primary key,
    chat_id int not null,
    action_time timestamp not null default now(),
    action varchar(50) not null
);

-- +goose Down
drop table chat_log;