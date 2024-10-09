-- +goose Up
create table chat
(
    id         serial primary key,
    chat_title text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table chat;