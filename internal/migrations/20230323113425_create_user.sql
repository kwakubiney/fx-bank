-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS "users"
(
    id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),
    username        VARCHAR(20) NOT NULL,
    password        VARCHAR(100) NOT NULL,
    created_at 		TIMESTAMP 	NOT NULL,
    last_modified   TIMESTAMP   NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "users";