-- +goose Up
-- +goose StatementBegin
create table if not exists program
(
    id             uuid primary key                        not null,
    title          varchar(500)                            not null,
    specialty_ugsn varchar(10) references specialty (code) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists program;
-- +goose StatementEnd
