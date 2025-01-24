-- +goose Up
-- +goose StatementBegin
create type transaction_type as enum ('credit', 'debit');

alter table operations_types
	add column transaction_type transaction_type;

update operations_types set transaction_type='debit' where operation_type_id=1;
update operations_types set transaction_type='debit' where operation_type_id=2;
update operations_types set transaction_type='debit' where operation_type_id=3;
update operations_types set transaction_type='credit' where operation_type_id=4;

alter table transactions
    add column balance DECIMAL(15,2);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type transaction_type;

alter table operations_types
	drop column transaction_type;

alter table transactions
    drop column balance;

-- +goose StatementEnd
