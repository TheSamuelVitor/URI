create table customers(
	id numeric primary key,
	name varchar,
	street varchar,
	city varchar,
	state char(2),
	credit_limit numeric
);


insert into customers(id, name, street, city, state, credit_limit) values(1, 'Pedro Augusto da Rocha', 'Rua Pedro Carlos Hoffman', 'Porto Alegre', 'RS', 700.00);
insert into customers(id, name, street, city, state, credit_limit) values(2, 'Antonio Carlos Mamel', 'Av. Pinheiros', 'Belo Horizonte', 'MG', 3500.50);
insert into customers(id, name, street, city, state, credit_limit) values(3, 'Luiza Augusta Mhor', 'Rua Salto Grande', 'Niteroi', 'RJ', 4000.00);
insert into customers(id, name, street, city, state, credit_limit) values(4, 'Jane Ester', 'Av 7 de setembro', 'Erechim', 'RS', 800.00);
insert into customers(id, name, street, city, state, credit_limit) values(5, 'Marcos Antônio dos Santos', 'Av Farrapos', 'Porto Alegre', 'RS', 4250.25);

select name from customers where state = 'RS'