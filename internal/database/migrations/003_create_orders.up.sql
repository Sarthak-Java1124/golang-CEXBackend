CREATE TABLE orders (
    id UUID PRIMARY KEY,

    market VARCHAR(50) NOT NULL,

    side VARCHAR(10) NOT NULL,
    type VARCHAR(20) NOT NULL,

    user_id UUID NOT NULL,

    status VARCHAR(20) NOT NULL,

    price NUMERIC(20,8) NOT NULL,

    quantity NUMERIC(20,8) NOT NULL,

    remaining_quantity NUMERIC(20,8) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);