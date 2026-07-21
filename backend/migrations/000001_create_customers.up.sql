CREATE EXTENSION IF NOT EXISTS "pgcrypto";


CREATE TABLE customers (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    first_name VARCHAR(100) NOT NULL,

    last_name VARCHAR(100) NOT NULL,

    email VARCHAR(255),

    phone VARCHAR(30),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMP NULL

);


CREATE INDEX idx_customers_email
ON customers(email);


CREATE INDEX idx_customers_deleted_at
ON customers(deleted_at);