-- +goose Up

ALTER TABLE specimen RENAME COLUMN collectionid TO collection_id;

-- +goose Down
