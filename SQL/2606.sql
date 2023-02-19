CREATE TABLE products
(
  id numeric PRIMARY KEY,
  name VARCHAR,
  amount NUMERIC,
  price numeric,
  id_categories numeric,
  FOREIGN KEY(id_categories)
  REFERENCES categories(id)
);

create table categories
(
  id numeric PRIMARY KEY,
  name VARCHAR
);


insert into products
  (id, name, amount, price, id_categories)
values(1, 'Lampshade', 100, 800, 4);
insert into products
  (id, name, amount, price, id_categories)
values(2, 'Table for painting', 1000, 560, 9);
insert into products
  (id, name, amount, price, id_categories)
values(3, 'Notebook desk', 10000, 25.50, 9);
insert into products
  (id, name, amount, price, id_categories)
values(4, 'Computer desk', 350, 320.50, 6);
insert into products
  (id, name, amount, price, id_categories)
values(5, 'Chair', 3000, 210.64, 9);
insert into products
  (id, name, amount, price, id_categories)
values(6, 'Home alarm', 750, 460, 4);

insert into categories
  (id, name)
values(1, 'old stock');
insert into categories
  (id, name)
values(2, 'new stock');
insert into categories
  (id, name)
values(3, 'modern');
insert into categories
  (id, name)
values(4, 'commercial');
insert into categories
  (id, name)
values(5, 'recyclable');
insert into categories
  (id, name)
values(6, 'executive');
insert into categories
  (id, name)
values(7, 'superior');
insert into categories
  (id, name)
values(8, 'wood');
insert into categories
  (id, name)
values(9, 'super luxury');
insert into categories
  (id, name)
values(10, 'vintage');