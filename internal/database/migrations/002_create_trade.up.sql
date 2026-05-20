CREATE TABLE trades (
    id VARCHAR(255) PRIMARY KEY,

    market VARCHAR(100) NOT NULL,

    buy_order_id VARCHAR(255) NOT NULL,
    sell_order_id VARCHAR(255) NOT NULL,

    buyer_user_id VARCHAR(255) NOT NULL,
    seller_user_id VARCHAR(255) NOT NULL,

    price BIGINT NOT NULL,
    quantity BIGINT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);