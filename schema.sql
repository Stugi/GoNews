DROP TABLE IF EXISTS authors, posts;

-- create a sequence
CREATE SEQUENCE id_seq START 1;

CREATE TABLE authors (
    id INTEGER PRIMARY KEY DEFAULT nextval ('id_seq'),
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY DEFAULT nextval ('id_seq'),
    author_id INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL,
    published_at BIGINT NOT NULL,
    FOREIGN KEY (author_id) REFERENCES authors (id)
);

INSERT INTO authors (name) VALUES ("author1");

INSERT authors (name) VALUES ('author2');