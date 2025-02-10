CREATE TABLE IF NOT EXISTS users (
 id UUID NOT NULL PRIMARY KEY,
 username VARCHAR(255) NOT NULL,
 hashed_password varchar(255) NOT NULL,
 coins INT NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(500) NOT NULL,
    price INT NOT NULL
);

CREATE TABLE IF NOT EXISTS coin_transactions(
    transaction_id UUID NOT NULL PRIMARY KEY,
    from_id UUID NOT NULL,
    to_id UUID NOT NULL,
    amount_of_coins INT NOT NULL
);

CREATE TABLE IF NOT EXISTS boughts(
    bought_id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    product_id INT NOT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

INSERT INTO products(product_name, price) VALUES
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500) ON CONFLICT DO NOTHING;

