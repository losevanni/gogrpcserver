CREATE TABLE users (
    idx SERIAL PRIMARY KEY,
    id TEXT NOT NULL,
    passwd TEXT NOT NULL
);

-- SERIAL PRIMARY KEY 자동 고유값 부여