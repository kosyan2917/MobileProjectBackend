CREATE TABLE pieces (
    id serial PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    filename VARCHAR(255) UNIQUE NOT NULL,
    length FLOAT
);