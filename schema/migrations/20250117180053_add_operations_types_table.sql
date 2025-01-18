-- +goose Up
-- +goose StatementBegin
CREATE TABLE operations_types (
    operation_type_id SERIAL PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS operations_types;
-- +goose StatementBegin
-- +goose StatementEnd