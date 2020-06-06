CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    password_hash varchar not null
);
