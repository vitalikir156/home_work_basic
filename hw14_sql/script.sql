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


-- напихаем таблицу с пользователями 
INSERT INTO market.users (name, email, password)
VALUES ('John Doe', 'john@example.com', crypt('some_password', gen_salt('md5'))),
('Sewa Arsch', 'SA@mymail.eu', crypt('JKjbzd!c87', gen_salt('md5')));
insert into market.users (name, password, email) values ('Leese Caro', '$2a$04$zkUEba4StBy8FCPKDYDz.eBFd75BQmiqXZ.h9ZiuXqr0ISOnj1xhG', 'lcaro0@dyndns.org');
insert into market.users (name, password, email) values ('Odille Meates', '$2a$04$EXuWLCbZsIDbQqYk3VlbAOsPi2jFjzdJr6J4Ruw3CFjlSsMsXBNsi', 'omeates1@discuz.net');
insert into market.users (name, password, email) values ('Oralee Dewdney', '$2a$04$d3HL4IN3y7NcfJ2VBxf5heus4.wV35kcqNyhwOM2d8LZWftoodo.q', 'odewdney2@sciencedaily.com');
insert into market.users (name, password, email) values ('Jo Wilby', '$2a$04$KhN9TCgIs7F5DZdNBYODTe/FX96X8YWrM/Nc/AmqW2VXHL8VmWdnK', 'jwilby3@dion.ne.jp');
insert into market.users (name, password, email) values ('Alfy Maides', '$2a$04$IxCwtfUWp0MBtQEF1yrmZOC7/r2Y/VoVKp9peVDRPykyg9cOZtgSC', 'amaides4@csmonitor.com');
insert into market.users (name, password, email) values ('Saba Pahler', '$2a$04$G.qfcAtUbBVrFDqj4jVzNu/CqP7nVPFWbau3uaJuc7hvEDt0IxWyu', 'spahler5@wufoo.com');
insert into market.users (name, password, email) values ('Bruno Kunat', '$2a$04$KGViKpBN2j8YvAwqOV243uLGnF1wt67mvJuCtOQJi2rkSkH.WZXo6', 'bkunat6@addthis.com');
insert into market.users (name, password, email) values ('Regina Illston', '$2a$04$EFgo/PzjMeE6GQ3Qr7CVR.XKZFkKTn1QhcQyZeGdKv2M3YImFwgPe', 'rillston7@alexa.com');
insert into market.users (name, password, email) values ('Lotte Garbert', '$2a$04$MJQVG.pF0I8xXpWakjgHQechn.f/VIU/4b.jCdkezG/Q8CjfldP5q', 'lgarbert8@rambler.ru');
insert into market.users (name, password, email) values ('Wilmer Newdick', '$2a$04$jkYqROJY4ga6S/6mpmfvHOStAIKOmZ5zYi8ZF7NwS5zpYRU2F6tlu', 'wnewdick9@networkadvertising.org');

-- напихаем таблицу с продуктами
INSERT INTO market.products (name, price) 
values ('ATTINY25-20SSUR, Микроконтроллер 2,7-5,5V 2K-Флэш-память 20МГц' , 199), 
('ATmega2560-16AU, Микроконтроллер 8-Бит, AVR, 16МГц, 256КБ Flash [TQFP-100]', 3000),
('ATMEGA3290P-20AU', 3310),
('ATMEGA1280-16CU, 8-bit Microcontrollers - MCU AVR 128K FLASH 8K SRAM 4KB EE-16 MHZ', 3950),
('ATMEGA324PB-AU, Микроконтроллер 8-бит ATmega 32 кБ Флэш-память 2 кБ Статическое ОЗУ 20 МГц TQFP-44', 410);

-- добавим пару заказов
insert into market.orders (user_id, order_date)
(SELECT users.id, '2077-11-11' from market.users where name='Lotte Garbert');
insert into market.orders (user_id, order_date)
(SELECT users.id, '2077-11-09' from market.users where name='Regina Illston');

-- напихаем в заказы товаров
INSERT INTO market.OrderProducts (orderId, ProductId) 
values (1, 1), (1, 2),(2, 4), (2, 3), (2, 5);

-- посчитаем суммы заказов и запихнём их в заказы
update market.orders 
SET total_amount = subqw.avg_price from (SELECT market.OrderProducts.orderId as oid, COALESCE(SUM(market.products.price), 0) AS avg_price FROM market.OrderProducts LEFT JOIN market.products ON products.id=productid GROUP BY oid) as subqw
where market.orders.id = subqw.oid;

-- Кто такой Jo Wilby? Удалим его
DELETE FROM market.users WHERE name = 'Jo Wilby';

-- А вот кто-то в мейле ошибся, исправляем
update market.users set email='bkunat@addthis.com' where name = 'Bruno Kunat';

-- Учитывая новые сборы необходимо поднять стоимость всех товаров которые дешевле 1000
update market.products set price=(price+price/10) where price < 1000;

-- А вот ATMEGA1280-16CU не продаётся больше. Удаляем. 
DELETE FROM market.products where id = 4;

-- Заказ 2 был отменён пользователем. Жаль.
delete from market.orders where id=2;

-- ещё 2 заказа!
insert into market.orders (user_id, order_date)
(SELECT users.id, '2077-11-11' from market.users where name='Lotte Garbert');
insert into market.orders (user_id, order_date)
(SELECT users.id, '2077-11-09' from market.users where name='Regina Illston');
INSERT INTO market.OrderProducts (orderId, ProductId) 
values (3, 1), (3, 2), (4, 3), (4, 5);

-- и ещё раз пересчитаем суммы заказов
update market.orders 
SET total_amount = subqw.avg_price from (SELECT market.OrderProducts.orderId as oid, COALESCE(SUM(market.products.price), 0) AS avg_price FROM market.OrderProducts LEFT JOIN market.products ON products.id=productid GROUP BY oid) as subqw
where market.orders.id = subqw.oid;

-- выборка пользователей (хотим посмотреть на всех у кого почта .com)
select id, name, email from market.users where email like '%.com';

-- выборка товаров (хотим увидеть какие меги у нас есть)
select * from market.products where lower(name) like lower('%atmega%');

-- Кто что поназаказывал? Включая товары
select user_id, users.name, order_date, orders.id, total_amount, price, products.name from ((market.orderproducts join market.orders on id=orderid) join market.products on products.id=productid) join market.users on users.id = user_id;

-- Заказы по пользакам 
select user_id, users.name, order_date, orders.id as "order id", total_amount from market.orders  join market.users on users.id = user_id;


-- средняя цена товара по каждому пользаку
SELECT subqw.uid, subqw.uname, COALESCE(AVG(subqw.prodprice), 0) AS avg_price
FROM (select user_id as uid, users.name as uname, order_date, orders.id, total_amount, price as prodprice, products.name from ((market.orderproducts join market.orders on id=orderid) join market.products on products.id=productid) join market.users on users.id = user_id) as subqw
GROUP BY subqw.uid, subqw.uname;

-- сумма товаров по каждому пользаку
SELECT subqw.uid, subqw.uname, COALESCE(SUM(subqw.total), 0) AS sum_price
FROM (select user_id as uid, users.name as uname, order_date, orders.id as "order id", total_amount as total from market.orders  join market.users on users.id = user_id) as subqw
GROUP BY subqw.uid, subqw.uname;

-- средняя цена товара в каждом заказе 
SELECT market.OrderProducts.orderId, COALESCE(AVG(market.products.price), 0) AS avg_price
FROM market.OrderProducts
LEFT JOIN market.products ON products.id=productid
GROUP BY market.OrderProducts.orderId;

-- сумма товара в каждом заказе 
SELECT market.OrderProducts.orderId, COALESCE(SUM(market.products.price), 0) AS avg_price
FROM market.OrderProducts
LEFT JOIN market.products ON products.id=productid
GROUP BY market.OrderProducts.orderId;

-- создадим индексы. 
create index on market.users(name);
create index on market.OrderProducts (orderid, productid);


