-- +goose Up
-- +goose StatementBegin
CREATE TABLE task_statuses(
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT,
    surname TEXT,
    login TEXT UNIQUE,
    password TEXT,
    email TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT,
    status INTEGER REFERENCES task_statuses(id), 
    author INTEGER REFERENCES users(id),
    executor INTEGER REFERENCES users(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE task_statuses;
DROP TABLE users;
DROP TABLE tasks;
-- +goose StatementEnd
