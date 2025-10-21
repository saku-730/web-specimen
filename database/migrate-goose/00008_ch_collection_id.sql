-- +goose Up

ALTER TABLE specimen ALTER COLUMN collection_id TYPE TEXT USING collection_id::TEXT;

-- +goose Down
