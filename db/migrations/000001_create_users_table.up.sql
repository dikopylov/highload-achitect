CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
    id         uuid primary key default uuid_generate_v4(),
    birthdate  date,
    first_name varchar not null,
    last_name  varchar not null,
    biography  varchar default '',
    city       varchar not null,
    password   varchar default '',
    age        int     not null check (age > 0)
);