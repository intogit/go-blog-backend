--  name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- // -- name: GetUsers :many
-- // select * from users

-- // -- name: GetUserById :one
-- // select * from users where id = $1
-- // -- name: DeleteUser :one
-- // -- name: UpdateUser :one
