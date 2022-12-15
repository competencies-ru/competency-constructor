-- +goose Up
-- +goose StatementBegin
alter table program drop column specialty_ugsn;
alter table program add column specialty_code varchar(10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table program drop column specialty_code;
alter table program add column specialty_ugsn varchar(10);
-- +goose StatementEnd
