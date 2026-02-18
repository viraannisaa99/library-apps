-- migrations/init.sql

CREATE TABLE IF NOT EXISTS authors (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL,
    bio        TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS books (
    id          SERIAL PRIMARY KEY,
    author_id   INT NOT NULL REFERENCES authors(id) ON DELETE CASCADE,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    published_year INT,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS reviews (
    id         SERIAL PRIMARY KEY,
    book_id    INT NOT NULL REFERENCES books(id) ON DELETE CASCADE,
    reviewer   VARCHAR(255) NOT NULL,
    rating     INT CHECK (rating BETWEEN 1 AND 5),
    comment    TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
