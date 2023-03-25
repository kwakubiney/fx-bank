-- +goose Up
CREATE TABLE IF NOT EXISTS transactions
(
    id 				UUID 		PRIMARY KEY DEFAULT gen_random_uuid(),
    credit          BIGINT      NOT NULL,
    debit           BIGINT      NOT NULL,
    created_at 		TIMESTAMP 	NOT NULL,
    provider_name   VARCHAR(50) NOT NULL,
    sender_currency VARCHAR(50) NOT NULL,
    receiver_currency VARCHAR(50) NOT NULL,
    sender_account_name VARCHAR(50)      NOT NULL,
    receiver_account_name VARCHAR(50) NOT NULL,
    status         VARCHAR(9) NOT NULL,
    updated_at      TIMESTAMP NOT NULL,
    rate            DECIMAL NOT NULL
    );

-- +goose Down
DROP TABLE IF EXISTS transactions;