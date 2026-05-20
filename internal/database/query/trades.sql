-- name: CreateTrade :one
INSERT INTO trades (
    id,
    market,
    buy_order_id,
    sell_order_id,
    buyer_user_id,
    seller_user_id,
    price,
    quantity,
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
    $9
)
RETURNING *;


-- name: GetTradeByID :one
SELECT *
FROM trades
WHERE id = $1;


-- name: ListTrades :many
SELECT *
FROM trades
ORDER BY created_at DESC;


-- name: GetTradesByMarket :many
SELECT *
FROM trades
WHERE market = $1
ORDER BY created_at DESC;


-- name: GetTradesByBuyer :many
SELECT *
FROM trades
WHERE buyer_user_id = $1
ORDER BY created_at DESC;


-- name: GetTradesBySeller :many
SELECT *
FROM trades
WHERE seller_user_id = $1
ORDER BY created_at DESC;


-- name: GetTradesByUser :many
SELECT *
FROM trades
WHERE buyer_user_id = $1
   OR seller_user_id = $1
ORDER BY created_at DESC;


-- name: GetRecentTradesByMarket :many
SELECT *
FROM trades
WHERE market = $1
ORDER BY created_at DESC
LIMIT $2;


-- name: GetMarketVolume :one
SELECT COALESCE(SUM(quantity), 0)
FROM trades
WHERE market = $1;


-- name: GetMarketTradeCount :one
SELECT COUNT(*)
FROM trades
WHERE market = $1;


-- name: DeleteTrade :exec
DELETE FROM trades
WHERE id = $1;