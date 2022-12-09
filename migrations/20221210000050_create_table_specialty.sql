-- +goose Up
-- +goose StatementBegin
create table if not exists specialty
(
    code      varchar(10) primary key            not null,
    title     varchar(500)                       not null,
    code_ugsn varchar(10) references ugsn (code) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists specialty;
-- +goose StatementEnd
