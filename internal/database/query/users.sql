-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    password,
    name,
    created_at
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;


-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;


-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;


-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;


-- name: UpdateUserName :exec
UPDATE users
SET name = $2
WHERE id = $1;


-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;