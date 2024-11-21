CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    customer_id VARCHAR(36) NOT NULL,
    total_amount FLOAT NOT NULL,
    payment_method VARCHAR(255) NOT NULL,
    payment_date TIMESTAMPTZ NOT NULL,
    payment_status VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);