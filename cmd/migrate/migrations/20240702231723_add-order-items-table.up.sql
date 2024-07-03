CREATE TABLE
    IF NOT EXISTS order_items (
        id SERIAL PRIMARY KEY,
        order_id INT NOT NULL,
        product_id INT NOT NULL,
        quantity INT NOT NULL,
        price DECIMAL(10, 2) NOT NULL,
        CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders (id),
        CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products (id)
    );