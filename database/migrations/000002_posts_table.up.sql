CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       content TEXT NOT NULL,
                       author_id INT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);
