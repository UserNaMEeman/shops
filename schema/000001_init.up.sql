CREATE TABLE users
(
    id serial not null unique,
    login varchar(255) not null unique,
    user_guid varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE orders
(
    id serial not null unique,
    user_guid varchar(255) not null unique,
    value varchar(255) not null
);