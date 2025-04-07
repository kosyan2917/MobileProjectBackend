CREATE TABLE tracks (
    id serial PRIMARY KEY,
    owner_id INT NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    time INT NOT NULL,
    created_at INT NOT NULL,
    distance FLOAT NOT NULL,
    CONSTRAINT fk_users FOREIGN KEY(owner_id) REFERENCES users(id)
);