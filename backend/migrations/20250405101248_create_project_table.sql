-- +goose Up
-- +goose StatementBegin
create table
    if not exists projects (
        id UUID primary key default uuid_generate_v4 (),
        name varchar(255) not null,
        description text,
        owner_id uuid not null references users (id) on delete cascade,
        status varchar(50) not null default 'active',
        created_at timestamp
        with
            time zone default now (),
            updated_at timestamp
        with
            time zone default now ()
    );

create index idx_projects_owner on projects (owner_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop index if exists idx_projects_owner;

drop table if exists projects;

-- +goose StatementEnd