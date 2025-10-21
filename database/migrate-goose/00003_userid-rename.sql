-- +goose Up

ALTER TABLE users RENAME COLUMN userid TO user_id;

-- +goose Down
