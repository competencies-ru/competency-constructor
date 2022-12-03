-- +goose Up
-- +goose StatementBegin
create table if not exists ugsn(
    code varchar(10) primary key not null,
    title varchar(500) not null
);


insert into ugsn(code, title) values ('', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists ugsn;
-- +goose StatementEnd
