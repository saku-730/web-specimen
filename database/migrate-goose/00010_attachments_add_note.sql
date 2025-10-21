-- +goose Up
ALTER TABLE attachments ADD COLUMN note TEXT;

-- +goose Down
