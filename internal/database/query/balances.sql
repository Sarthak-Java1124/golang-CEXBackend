-- name: CreateBalance :one
INSERT INTO balances (
    id,
    user_id,
    asset,
    price,
    quantity
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;


-- name: GetBalanceByID :one
SELECT *
FROM balances
WHERE id = $1;


-- name: GetUserBalances :many
SELECT *
FROM balances
WHERE user_id = $1
ORDER BY asset ASC;


-- name: GetUserAssetBalance :one
SELECT *
FROM balances
WHERE user_id = $1
  AND asset = $2;


-- name: UpdateBalancePrice :exec
UPDATE balances
SET price = $3
WHERE user_id = $1
  AND asset = $2;


-- name: IncreaseBalanceQuantity :exec
UPDATE balances
SET quantity = quantity + $3
WHERE user_id = $1
  AND asset = $2;


-- name: DecreaseBalanceQuantity :exec
UPDATE balances
SET quantity = quantity - $3
WHERE user_id = $1
  AND asset = $2;


-- name: DeleteBalance :exec
DELETE FROM balances
WHERE id = $1;