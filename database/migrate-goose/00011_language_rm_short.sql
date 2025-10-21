-- +goose Up
ALTER TABLE languages DROP COLUMN language_short;

-- +goose Down
