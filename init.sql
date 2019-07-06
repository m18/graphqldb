create table customers (
    id serial primary key,
    name varchar not null
);

insert into customers(name) values ('andrew');
insert into customers(name) values ('cathy');
insert into customers(name) values ('max');
insert into customers(name) values ('jackie');

create table products (
    id serial primary key,
    name varchar not null,
    price integer not null -- in cents
);

insert into products (name, price) values ('toothbrush', 399);
insert into products (name, price) values ('toothpaste', 349);
insert into products (name, price) values ('soap', 249);
insert into products (name, price) values ('face wash', 899);

create table orders (
    id serial primary key,
    customer_id integer references customers(id),
    time timestamptz not null
);

insert into orders(customer_id, time) values (1, now() + (-2 * interval '1 day'));
insert into orders(customer_id, time) values (1, now() + (-8 * interval '1 day'));
insert into orders(customer_id, time) values (2, now() + (-1 * interval '1 day'));
insert into orders(customer_id, time) values (3, now() + (-5 * interval '1 day'));

create table order_products (
    id serial primary key,
    order_id integer references orders(id),
    product_id integer references products(id),
    quantity integer not null
);

insert into order_products(order_id, product_id, quantity) values (1, 1, 2);
insert into order_products(order_id, product_id, quantity) values (1, 2, 1);
insert into order_products(order_id, product_id, quantity) values (2, 3, 4);
insert into order_products(order_id, product_id, quantity) values (3, 3, 2);
insert into order_products(order_id, product_id, quantity) values (3, 4, 1);
insert into order_products(order_id, product_id, quantity) values (4, 4, 2);
insert into order_products(order_id, product_id, quantity) values (4, 2, 2);
