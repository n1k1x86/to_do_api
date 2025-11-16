-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY,
    login TEXT,
    password TEXT,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT login_unique UNIQUE (login)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table users;
-- +goose StatementEnd
