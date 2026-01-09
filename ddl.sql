-- DDL
CREATE TABLE user(
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL, 
    first_name VARCHAR(100) NOT NULL, 
    last_name VARCHAR(100) NOT NULL, 
    password TEXT NOT NULL,
    profile_image TEXT
);

CREATE TABLE banner(
    id SERIAL PRIMARY KEY,
    banner_name VARCHAR(100) NOT NULL,
    banner_image TEXT NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE service(
    id SERIAL PRIMARY KEY,
    service_code VARCHAR(100) NOT NULL,
    service_icon TEXT NOT NULL,
    service_tariff DECIMAL(15,2) NOT NULL CHECK (service_tariff > 0)
);

CREATE TYPE transaction_status AS ENUM ('pending', 'shipped', 'cancelled');

CREATE TABLE transaction(
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(100) REFERENCES user(email),
    invoice_number VARCHAR(100) NOT NULL,
    service_code VARCHAR(100) NOT NULL,
    service_name VARCHAR(100) NOT NULL,
    transaction_type transaction_status,
    total_amount DECIMAL (15,2) NOT NULL,
    created_on TIMESTAMP DEFAULT NOW()
);