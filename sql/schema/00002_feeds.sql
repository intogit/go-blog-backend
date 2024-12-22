-- +goose Up
-- +goose StatementBegin
CREATE TABLE feeds (
    feed_id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    feed_name TEXT NOT NULL,
    feed_url text unique not null,
    user_id UUID NOT NULL references users(user_id) on delete cascade
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds;
-- +goose StatementEnd
