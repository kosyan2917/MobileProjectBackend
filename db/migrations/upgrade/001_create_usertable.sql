CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    avatar VARCHAR(255),
    name VARCHAR(255)
);