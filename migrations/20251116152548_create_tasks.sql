-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks(
    id INTEGER PRIMARY KEY,
    title TEXT,
    description TEXT,
    author INTEGER,
    estimation INTEGER,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT author_fk FOREIGN KEY (author) REFERENCES users(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tasks;
-- +goose StatementEnd
