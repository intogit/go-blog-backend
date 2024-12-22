-- name: CreateFeedFollow :one
INSERT INTO feed_follows (feed_follows_id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllFeedFollow :many
SELECT * from feed_follows where user_id = $1;

-- name: DeleteFeedFollow :exec
delete from feed_follows where feed_follows_id = $1 and user_id = $2;
-- // -- reason to use user_id along with feed_follows_id is that
-- // -- what if user B get feed_follows_id of A anyhow, then B will be able to delete the A feed feed_follows
-- // -- so, if we have check that A is only deleting eed_follows_id of A only.