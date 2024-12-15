CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    token VARCHAR(255) default null,
    started timestamp default null
);