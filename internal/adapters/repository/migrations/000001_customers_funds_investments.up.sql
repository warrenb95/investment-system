CREATE TABLE customers (
    id UUID PRIMARY KEY
);

CREATE TABLE funds (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE investments (
    id UUID PRIMARY KEY,
    customer_id UUID REFERENCES customers(id) ON DELETE CASCADE,
    fund_id UUID REFERENCES funds(id) ON DELETE CASCADE,
    amount DECIMAL NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP DEFAULT NOW()
);
