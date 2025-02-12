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

INSERT INTO users(id, username, hashed_password, coins) VALUES
    ('2641b07b-ef83-4eeb-9734-71e78248cd5f', 'user2', '$2a$07$dpqSk3DZtNCab5pSQKeyWujy19Lso2usHOdVHh0NyFzoY60E/S/Y.', 388),
    ('3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb', 'user1', '$2a$07$yLZuAQlDhsBhkrVTO/bOfeUEF.2xFpYDMFAN9KSDOS9XM5ZYY24cK', 615),
    ('fdae7770-ee51-4563-8261-704beb45478d', 'user3', '$2a$07$rJ5ixYA62XpdLgdURHFK3.lkLOvCcR0CZWdsnIw7MC3Oj1IY0cOgK', 667);

INSERT INTO boughts(bought_id, user_id, product_id) VALUES
    ('1932117f-5145-4871-af73-60066665a17a','fdae7770-ee51-4563-8261-704beb45478d',6),
    ('1c3d07d0-86ab-460a-b28e-dde642efd35b','2641b07b-ef83-4eeb-9734-71e78248cd5f',7),
    ('1cff9b31-3eb4-4781-b02b-6c27ba63e4b6','fdae7770-ee51-4563-8261-704beb45478d',9),
    ('7b82138e-cf55-4762-8584-d4381838a2a4','2641b07b-ef83-4eeb-9734-71e78248cd5f',1),
    ('a08849f5-613e-4f49-aa5d-22418692c35d','fdae7770-ee51-4563-8261-704beb45478d',6),
    ('b211a891-5e78-45ae-8610-6f69359c0664','3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb',8),
    ('d0f6b290-e798-49f8-a581-e123b2acf6b6','2641b07b-ef83-4eeb-9734-71e78248cd5f',1),
    ('da5336d8-b137-4cdd-91f6-8d8403116417','3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb',6),
    ('edd16979-33d1-48c8-aeff-4761977cca3d','3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb',8);

INSERT INTO coin_transactions (transaction_id, from_id, to_id, amount_of_coins) VALUES
    ('0ebbe4d6-4aa7-4cd6-afe4-b37ab4ba2a50', '2641b07b-ef83-4eeb-9734-71e78248cd5f', '3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb', 200),
    ('2d3bb251-0a32-4fb6-8c4f-c1542c0eddd1', '2641b07b-ef83-4eeb-9734-71e78248cd5f', 'fdae7770-ee51-4563-8261-704beb45478d', 112),
    ('2d900b8a-3623-4fe4-b064-1b57c50cca04', '2641b07b-ef83-4eeb-9734-71e78248cd5f', 'fdae7770-ee51-4563-8261-704beb45478d', 45),
    ('3eee3a3e-4b05-46f0-935f-cbea41de9ef4', 'fdae7770-ee51-4563-8261-704beb45478d', '2641b07b-ef83-4eeb-9734-71e78248cd5f', 30),
    ('92eff357-dc88-44cf-b0e8-d554ade69ccc', 'fdae7770-ee51-4563-8261-704beb45478d', '2641b07b-ef83-4eeb-9734-71e78248cd5f', 10),
    ('a57fd7f0-587f-46d3-a16e-261c0d70d6b1', '3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb', 'fdae7770-ee51-4563-8261-704beb45478d', 400),
    ('aaf07e70-48b1-4040-ab32-1bdd8cabf894', '3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb', '2641b07b-ef83-4eeb-9734-71e78248cd5f', 35),
    ('b573a458-76a7-44ea-bdfc-703c58ddd343', 'fdae7770-ee51-4563-8261-704beb45478d', '3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb', 50),
    ('ed62c5e1-bc40-4980-87db-7c2ccdc3fd5a', 'fdae7770-ee51-4563-8261-704beb45478d', '3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb', 150),
    ('ffe60270-1116-4593-9f61-fa8d907206a4', '3614bbf5-01ad-4a86-a9cb-cc0fbebda6fb', '2641b07b-ef83-4eeb-9734-71e78248cd5f', 30);

