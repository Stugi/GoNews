DROP TABLE IF EXISTS authors, posts;

CREATE TABLE authors (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    author_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    published_at BIGINT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES authors (id)
);