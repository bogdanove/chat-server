-- +goose Up
create table message
(
    id        serial primary key,
    msg_from      text      not null,
    msg_text      text      not null,
    msg_timestamp timestamp not null default now()
);

-- +goose Down
drop table message;
