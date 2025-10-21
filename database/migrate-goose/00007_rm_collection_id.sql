-- +goose Up

ALTER TABLE specimen DROP CONSTRAINT specimen_collectionid_fkey;
DROP TABLE collection_id_code

-- +goose Down
