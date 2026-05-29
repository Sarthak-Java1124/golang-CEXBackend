-- name: CreateBalance :one
INSERT INTO balances (
    id,
    user_id,
    balance,
    asset_balance
)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: GetBalanceByID :one
SELECT *
FROM balances
WHERE id = $1;


-- name: GetBalanceByUserID :one
SELECT *
FROM balances
WHERE user_id = $1;


-- name: ListBalances :many
SELECT *
FROM balances
ORDER BY user_id;


-- name: UpdateBalance :one
UPDATE balances
SET
    balance = $2,
    asset_balance = $3
WHERE id = $1
RETURNING *;


-- name: UpdateAssetBalance :one
UPDATE balances
SET asset_balance = $2
WHERE id = $1
RETURNING *;


-- name: UpdateBalanceAmount :one
UPDATE balances
SET balance = $2
WHERE id = $1
RETURNING *;


-- name: DeleteBalance :exec
DELETE FROM balances
WHERE id = $1;


-- name: DeleteBalanceByUserID :exec
DELETE FROM balances
WHERE user_id = $1;