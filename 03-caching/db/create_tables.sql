CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

INSERT INTO product (name, price) VALUES ('Product 1', 15000);
INSERT INTO product (name, price) VALUES ('Product 2', 25000);
INSERT INTO product (name, price) VALUES ('Product 3', 35000);