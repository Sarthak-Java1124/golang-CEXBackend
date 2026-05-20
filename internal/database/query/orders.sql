-- name: CreateOrder :one
INSERT INTO orders (
    id,
    market,
    side,
    type,
    user_id,
    status,
    price,
    quantity,
    remaining_quantity,
    created_at
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
)
RETURNING *;


-- name: GetOrderByID :one
SELECT *
FROM orders
WHERE id = $1;


-- name: ListOrders :many
SELECT *
FROM orders
ORDER BY created_at DESC;


-- name: GetOrdersByUser :many
SELECT *
FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;


-- name: GetOrdersByMarket :many
SELECT *
FROM orders
WHERE market = $1
ORDER BY created_at DESC;


-- name: GetOpenBuyOrders :many
SELECT *
FROM orders
WHERE market = $1
  AND side = 'BUY'
  AND status = 'OPEN'
ORDER BY price DESC, created_at ASC;


-- name: GetOpenSellOrders :many
SELECT *
FROM orders
WHERE market = $1
  AND side = 'SELL'
  AND status = 'OPEN'
ORDER BY price ASC, created_at ASC;


-- name: GetMatchingSellOrders :many
SELECT *
FROM orders
WHERE market = $1
  AND side = 'SELL'
  AND status = 'OPEN'
  AND price <= $2
ORDER BY price ASC, created_at ASC;


-- name: GetMatchingBuyOrders :many
SELECT *
FROM orders
WHERE market = $1
  AND side = 'BUY'
  AND status = 'OPEN'
  AND price >= $2
ORDER BY price DESC, created_at ASC;


-- name: UpdateOrderRemainingQuantity :exec
UPDATE orders
SET remaining_quantity = $2
WHERE id = $1;


-- name: MarkOrderPartiallyFilled :exec
UPDATE orders
SET status = 'PARTIALLY_FILLED',
    remaining_quantity = $2
WHERE id = $1;


-- name: MarkOrderFilled :exec
UPDATE orders
SET status = 'FILLED',
    remaining_quantity = 0
WHERE id = $1;


-- name: CancelOrder :exec
UPDATE orders
SET status = 'CANCELLED'
WHERE id = $1;


-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;