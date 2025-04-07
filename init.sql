CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255)
);

INSERT INTO users (username, password)
VALUES
    ('test1', 'test'),
    ('test2', 'test'),
    ('test3', 'test');

CREATE TABLE tracks (
    id serial PRIMARY KEY,
    owner_id INT NOT NULL,
    name VARCHAR(255) UNIQUE NOT NULL,
    time INT NOT NULL,
    created_at INT NOT NULL,
    distance FLOAT NOT NULL,
    CONSTRAINT fk_users FOREIGN KEY(owner_id) REFERENCES users(id)
);

INSERT INTO tracks (owner_id, name, time, created_at, distance)
VALUES
    (1, 'Дневной_велозаезд (1)', 1231, 1743717685, 5),
    (1, 'Дневной_велозаезд (2)', 3213, 1743717685, 10),
    (1, 'Дневной_велозаезд (3)', 311, 1743717685, 7),
    (1, 'Дневной_велозаезд', 111, 1743717685, 8),
    (2, 'Дневной_велозаезд (4)', 332, 1743717685, 6),
    (2, 'Дневной_велозаезд (5)', 555, 1743717685, 5),
    (3, 'Дневной_велозаезд (6)', 1000, 1743717685, 3),
    (3, 'Послеобеденный_велозаезд', 2000, 1743717685, 6);