-- +goose Up
ALTER TABLE users ADD COLUMN username TEXT;

-- +goose Down
SELECT 'down SQL query';
