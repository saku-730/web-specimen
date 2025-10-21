-- +goose Up

ALTER TABLE attachment_goup RENAME TO attachment_group;

-- +goose Down
