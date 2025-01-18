-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts (
    account_id SERIAL PRIMARY KEY,
    document_number VARCHAR(20) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS accounts;
-- +goose StatementBegin
-- +goose StatementEnd