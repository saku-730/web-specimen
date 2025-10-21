-- +goose Up

ALTER TABLE occurrence ALTER COLUMN timezone DROP NOT NULL;
ALTER TABLE observations ALTER COLUMN timezone DROP NOT NULL;
ALTER TABLE make_specimen ALTER COLUMN timezone DROP NOT NULL;
ALTER TABLE identifications ALTER COLUMN timezone DROP NOT NULL;
ALTER TABLE users ALTER COLUMN timezone DROP NOT NULL;

-- +goose Down
