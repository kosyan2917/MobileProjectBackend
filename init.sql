CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    avatar VARCHAR(255),
    name VARCHAR(255)

);

INSERT INTO users (username, password, name, avatar)
VALUES
    ('test1', 'test', 'Тестовый пользователь1', ''),
    ('test2', 'test', 'Тестовый пользователь2', ''),
    ('test3', 'test', 'Тестовый пользователь3', 'avatar.jpg');

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

CREATE TABLE pieces (
    id serial PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    filename VARCHAR(255) UNIQUE NOT NULL,
    length FLOAT
);

INSERT INTO pieces (name, filename, length)
VALUES
    ('Дорога в Новогиреево', 'piece.gpx', 0.71),
    ('Обратная дорога в Новогиреево', 'piecereversed.gpx', 0.71);

CREATE TABLE subscribed_pieces (
    id serial PRIMARY KEY,
    piece_id INT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_piece FOREIGN KEY(piece_id) REFERENCES pieces(id)
);

INSERT INTO subscribed_pieces (piece_id, user_id)
VALUES
    (1, 1),
    (2, 1);
