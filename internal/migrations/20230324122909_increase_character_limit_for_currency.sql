-- +goose Up
ALTER TABLE "accounts" ALTER COLUMN currency TYPE VARCHAR(100);
-- +goose Down
ALTER TABLE "accounts" ALTER COLUMN currency TYPE VARCHAR(15);