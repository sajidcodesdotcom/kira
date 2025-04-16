-- +goose Up
-- +goose StatementBegin
create table
    if not exists tasks (
        id uuid primary key default uuid_generate_v4 (),
        title varchar(255) not null,
        description text,
        project_id uuid not null references projects (id) on delete cascade,
        assignee_id uuid references users (id) on delete set null,
        status varchar(50) not null default 'todo',
        priority varchar(50) not null default 'medium',
        due_date timestamp
        with
            time zone,
            created_at timestamp
        with
            time zone default now (),
            updated_at timestamp
        with
            time zone default now ()
    );

create index idx_tasks_project on tasks (project_id);

create index idx_tasks_assignee on tasks (assignee_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop index if exists idx_tasks_project;

drop index if exists idx_tasks_assignee;

drop table if exists tasks;

-- +goose StatementEnd