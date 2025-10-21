-- +goose Up

ALTER TABLE language RENAME TO languages;

-- +goose Down
