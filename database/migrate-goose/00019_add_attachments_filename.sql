-- +goose Up

ALTER TABLE public.attachments ADD COLUMN original_filename text;

-- +goose Down
