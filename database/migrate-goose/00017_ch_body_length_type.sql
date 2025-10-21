-- +goose Up

ALTER TABLE occurrence ALTER COLUMN body_length TYPE text;

-- +goose Down
