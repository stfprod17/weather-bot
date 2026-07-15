-- +goose Up
CREATE TABLE users (
id integer primary key,
city text,
created_at timestamp default NOW()
);
-- +goose Down

