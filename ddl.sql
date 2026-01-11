-- DDL

CREATE DATABASE nutech;

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL, 
    first_name VARCHAR(100) NOT NULL, 
    last_name VARCHAR(100) NOT NULL, 
    password TEXT NOT NULL,
    profile_image TEXT DEFAULT ''
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
    service_name VARCHAR(100) NOT NULL,
    service_icon TEXT NOT NULL,
    service_tariff DECIMAL(15,2) NOT NULL CHECK (service_tariff > 0)
);

CREATE TYPE transaction_status AS ENUM ('pending', 'shipped', 'cancelled');

CREATE TABLE transaction(
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(100) REFERENCES users(email) ON DELETE CASCADE,
    invoice_number VARCHAR(100) NOT NULL,
    service_code VARCHAR(100) NOT NULL,
    service_name VARCHAR(100) NOT NULL,
    transaction_type transaction_status,
    total_amount DECIMAL (15,2) NOT NULL,
    created_on TIMESTAMP DEFAULT NOW()
);

CREATE TABLE saldo(
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(100) REFERENCES users(email) ON DELETE CASCADE,
    balance DECIMAL (15,2) DEFAULT 0 NOT NULL
);

INSERT INTO service(service_code, service_name, service_icon, service_tariff) 
VALUES
('PAJAK', 'PAJAK PBB', 'https://nutech-integrasi.app/dummy.jpg', 40000),
('PLN', 'Listrik', 'https://nutech-integrasi.app/dummy.jpg', 10000),
('PDAM', 'PDAM Berlangganan', 'https://nutech-integrasi.app/dummy.jpg', 40000),
('PULSA', 'Pulsa', 'https://nutech-integrasi.app/dummy.jpg', 40000),
('PGN', 'PGN Berlangganan', 'https://nutech-integrasi.app/dummy.jpg', 50000),
('MUSIK', 'MUSIK Berlangganan', 'https://nutech-integrasi.app/dummy.jpg', 50000),
('TV', 'TV Berlangganan', 'https://nutech-integrasi.app/dummy.jpg', 50000),
('PAKET DATA', 'Paket data', 'https://nutech-integrasi.app/dummy.jpg', 50000),
('VOUCHER MAKANAN', 'Voucher Makanan', 'https://nutech-integrasi.app/dummy.jpg', 100000),
('QURBAN', 'Qurban', 'https://nutech-integrasi.app/dummy.jpg', 200000),
('ZAKAT', 'Zakat', 'https://nutech-integrasi.app/dummy.jpg', 300000),
('VOUCHER GAME', 'Voucher Game', 'https://nutech-integrasi.app/dummy.jpg', 100000);
   