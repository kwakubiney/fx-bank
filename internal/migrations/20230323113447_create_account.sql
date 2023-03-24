
-- +goose Up
CREATE TABLE IF NOT EXISTS "accounts"
(
    id 		        UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(50) NOT NULL,
    balance         DECIMAL     NOT NULL,
    created_at 	    TIMESTAMP 	NOT NULL,
    last_modified   TIMESTAMP   NOT NULL,
    currency        VARCHAR(10) NOT NULL,
    user_id         UUID        NOT NULL,
    CONSTRAINT FK_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);
-- +goose Down
DROP TABLE IF EXISTS accounts;