CREATE TABLE vehicles (

    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    customer_id UUID NOT NULL,

    vin VARCHAR(17),

    year INT NOT NULL,

    make VARCHAR(100) NOT NULL,

    model VARCHAR(100) NOT NULL,

    trim VARCHAR(100),

    engine VARCHAR(100),

    drivetrain VARCHAR(50),

    transmission VARCHAR(50),

    mileage INT DEFAULT 0,

    color VARCHAR(50),

    license_plate VARCHAR(30),

    state VARCHAR(20),

    is_primary BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_vehicle_customer
        FOREIGN KEY (customer_id)
        REFERENCES customer_profiles(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_vehicle_customer
ON vehicles(customer_id);

CREATE INDEX idx_vehicle_vin
ON vehicles(vin);

CREATE INDEX idx_vehicle_primary
ON vehicles(customer_id, is_primary);