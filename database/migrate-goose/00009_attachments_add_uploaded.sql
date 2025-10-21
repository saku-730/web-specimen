-- +goose Up

ALTER TABLE attachments ADD COLUMN uploaded TIMESTAMP WITH TIME ZONE;

-- +goose Down
