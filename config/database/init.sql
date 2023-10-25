CREATE DATABASE laundry_api


CREATE TABLE product (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    price BIGINT,
    uom VARCHAR(10)
);

CREATE TABLE customer (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100),
    phone_number VARCHAR(20) UNIQUE,
    address TEXT
);

CREATE TABLE tx_bill (
    id VARCHAR(100) PRIMARY KEY,
    bill_date DATE,
    entry_date DATE,
    customer_id VARCHAR(100),
    CONSTRAINT fk_customer_id FOREIGN KEY(customer_id) REFERENCES customer(id)
);

CREATE TABLE tx_bill_detail (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    bill_id VARCHAR(100),
    customer_id VARCHAR(100),
    customer_price BIGINT,
    qty INT,
    finish_date DATE,
    CONSTRAINT fk_bill_id FOREIGN KEY(bill_id) REFERENCES tx_bill(id),
    CONSTRAINT fk_service_id FOREIGN KEY(customer_id) REFERENCES customer(id)
);

CREATE TABLE user_credential ( 
    id varchar(100) PRIMARY KEY not null,
    username varchar(100) UNIQUE not null,
    password varchar(200) not null,
    is_active BOOLEAN DEFAULT true
);
