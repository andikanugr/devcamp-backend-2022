CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    price NUMERIC NOT NULL
);

INSERT INTO product (name, description, price) VALUES ('Product 1', 'description of product 1', 15000);
INSERT INTO product (name, description, price) VALUES ('Product 2', 'description of product 1', 25000);
INSERT INTO product (name, description, price) VALUES ('Product 3', 'description of product 1', 35000);