-- migrate:up
create table if not exists "authors" (
    id bigserial primary key,
    "name" text not null,
    bio text not null,
    birth_date timestamp not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz
);
create index authors_id_idx on "authors" (id);

-- migrate:down
drop table if exists "authors";
