-- +goose Up
create table chat_users
(
    id      serial primary key,
    chat_id int not null references chat(id),
    user_id int not null
);

-- +goose Down
drop table chat_users;