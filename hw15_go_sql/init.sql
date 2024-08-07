-- шифратор паролей
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- генератор UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- создаём схему 'market'
create schema market;

-- далее создаём нужные таблицы
-- таблица пользаков
create table market.users (
id uuid PRIMARY KEY default uuid_generate_v4(),
name VARCHAR(255),
email VARCHAR(255),
password varchar(255));


-- таблица заказов
create table market.orders (
id serial PRIMARY KEY,
user_id uuid references market.users(id) ON DELETE CASCADE,
order_date DATE NOT NULL,
total_amount decimal);

-- таблица товаров
create table market.products (
id serial PRIMARY KEY,
name VARCHAR(255),
price decimal);

-- таблица заказы-товары
create table market.OrderProducts (
orderId integer references market.orders(id) ON DELETE CASCADE,
ProductId integer  references market.products(id) ON DELETE CASCADE,
primary key (orderId, ProductId));
