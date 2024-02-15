CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    refresh_token VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);