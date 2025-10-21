-- +goose Up

ALTER TABLE public.occurrence ALTER COLUMN timezone TYPE text;

ALTER TABLE public.observations ALTER COLUMN timezone TYPE text;

ALTER TABLE public.make_specimen ALTER COLUMN timezone TYPE text;

ALTER TABLE public.identifications ALTER COLUMN timezone TYPE text;

ALTER TABLE public.users ALTER COLUMN timezone TYPE text;

-- +goose Down
