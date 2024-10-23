-- +goose Up
create table chat
(
    id         serial primary key,
    chat_title varchar      not null,
    created_at timestamp not null default now()
);

-- +goose Down
drop table chat;