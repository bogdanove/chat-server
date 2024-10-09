-- +goose Up
create table chat_users
(
    id         serial primary key,
    chat_id    int       not null,
    user_id    int       not null
);

-- +goose Down
drop table chat_users;