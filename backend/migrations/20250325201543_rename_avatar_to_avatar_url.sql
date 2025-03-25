-- +goose Up
-- +goose StatementBegin
alter table users
rename column avatar to avatar_url;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
alter table users rename avatar_url to avatar;

-- +goose StatementEnd