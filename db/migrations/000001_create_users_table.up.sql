CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table users
(
    id          uuid primary key default uuid_generate_v4(),
    birthdate   date    not null,
    first_name  varchar not null,
    second_name varchar not null,
    biography   varchar not null,
    city        varchar not null,
    password    varchar not null,
    age         int     not null check (age > 0)
);