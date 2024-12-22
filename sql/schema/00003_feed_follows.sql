-- +goose Up
-- +goose StatementBegin
CREATE TABLE feed_follows (
    feed_follows_id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL references users(user_id) on delete cascade,
    feed_id UUID NOT NULL references feeds(feed_id) on delete cascade,
    UNIQUE(user_id, feed_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table feed_follows;
-- +goose StatementEnd
