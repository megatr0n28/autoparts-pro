CREATE TABLE customer_profiles (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL UNIQUE,

    first_name VARCHAR(100) NOT NULL,

    last_name VARCHAR(100) NOT NULL,

    phone VARCHAR(30),

    address_line1 VARCHAR(255),

    address_line2 VARCHAR(255),

    city VARCHAR(100),

    state VARCHAR(50),

    postal_code VARCHAR(20),

    country VARCHAR(50)
        DEFAULT 'USA',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),


    CONSTRAINT fk_customer_profiles_user

        FOREIGN KEY(user_id)

        REFERENCES users(id)

        ON DELETE CASCADE
);


CREATE INDEX idx_customer_profiles_user_id
ON customer_profiles(user_id);