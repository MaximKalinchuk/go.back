CREATE TABLE users (
    id UUID primary key not null,
    username varchar(80) not null,
    email varchar(255) not null unique,
    password_hash varchar(255) not null
);