-- name: CreateUser :one
INSERT INTO users (user_id, created_at, updated_at, user_name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByApiKey :one
select * from users where api_key = $1;

-- // -- name: GetUsers :many
-- // select * from users

-- // -- name: DeleteUser :one
-- // -- name: UpdateUser :one
