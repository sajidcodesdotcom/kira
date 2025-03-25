-- +goose Up
-- +goose StatementBegin
create extension if not exists "uuid-ossp";

create table
    if not exists users (
        id uuid primary key default uuid_generate_v4 (),
        email varchar(255) unique not NULL,
        username varchar(50) unique not null,
        full_name varchar(100) not null,
        password varchar(255) not null,
        avatar varchar(255),
        created_at timestamp
        with
            time zone default now (),
            updated_at timestamp
        with
            time zone default now ()
    );

create index idx_users_email on users (email);

create index idx_users_username on users (username);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop index if exists idx_users_email;

drop index if exists idx_users_username;

drop table if exists users;

-- +goose StatementEnd