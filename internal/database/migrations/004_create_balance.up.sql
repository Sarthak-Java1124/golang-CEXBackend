CREATE TABLE balances (
    id UUID PRIMARY KEY,

    user_id UUID NOT NULL,

    asset VARCHAR(20) NOT NULL,

    price NUMERIC(20,8) NOT NULL
        CHECK (price >= 0),

    quantity NUMERIC(20,8) NOT NULL
        CHECK (quantity >= 0)
);