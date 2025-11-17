-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY NOT NULL,
    login TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    CONSTRAINT login_unique UNIQUE (login)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table users;
-- +goose StatementEnd
