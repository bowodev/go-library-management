-- migrate:up
create table if not exists "books" (
    id bigserial primary key,
    title text not null,
    "description" text not null,
    publish_date timestamp not null,
    author_id bigint not null,
    created_at timestamptz default now(),
    updated_at timestamptz default now(),
    deleted_at timestamptz,
    CONSTRAINT fk_author FOREIGN KEY (author_id)
    REFERENCES authors (id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);
create index books_id_idx on "books" (id);
create index books_author_id_idx on "books" (author_id);

-- migrate:down
drop table if exists "books";

