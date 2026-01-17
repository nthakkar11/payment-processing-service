-- Customers table
CREATE TABLE customers (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Payment intents
CREATE TABLE payment_intents (
    id UUID PRIMARY KEY,
    customer_id UUID NOT NULL REFERENCES customers(id),
    amount BIGINT NOT NULL,
    currency TEXT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Transactions
CREATE TABLE transactions (
    id UUID PRIMARY KEY,
    payment_intent_id UUID NOT NULL REFERENCES payment_intents(id),
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Idempotency keys
CREATE TABLE idempotency_keys (
    key TEXT PRIMARY KEY,
    response JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
