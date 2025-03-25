-- +goose Up
-- +goose StatementBegin
alter table users
add column role varchar(50) not null default 'user';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
alter table users
drop column role;

-- +goose StatementEnd