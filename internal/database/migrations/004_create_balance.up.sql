CREATE TABLE balances (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    balance INTEGER NOT NULL DEFAULT 0,
    asset_balance JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);