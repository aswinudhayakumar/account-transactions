-- +goose Up
-- +goose StatementBegin
-- A seeder to add data in operations_types table
INSERT INTO operations_types (description) VALUES 
	('Normal Purchase'),
	('Purchase With Installments'),
	('Withdrawal'),
	('Credit Voucher');
-- +goose StatementEnd

-- +goose Down
DELETE FROM operations_types WHERE description IN 
    ('Normal Purchase', 'Purchase With Installments', 'Withdrawal', 'Credit Voucher');
-- +goose StatementBegin
-- +goose StatementEnd