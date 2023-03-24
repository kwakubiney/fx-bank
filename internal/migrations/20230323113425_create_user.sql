-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS "users"
(
    id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name      VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50) NOT NULL,
    created_at 		TIMESTAMP 	NOT NULL,
    last_modified   TIMESTAMP   NOT NULL,
    email           VARCHAR(50)  NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "users";