-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_name TEXT NOT NULL,
    api_key varchar(64) unique not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
