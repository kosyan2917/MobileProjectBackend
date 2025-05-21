CREATE TABLE subscribed_pieces (
    id serial PRIMARY KEY,
    piece_id INT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_piece FOREIGN KEY(piece_id) REFERENCES pieces(id)
);