create table users (
	id serial not null,
	name varchar(25) not null,
	email varchar(255) not null unique,
	phoneno varchar(25) not null unique,
	role int not null,
	primary key(id)
);
create table products (
	id serial not null,
	name varchar(25) not null ,
	price real not null ,
	tax real not null,
	seller_id int not null,
	primary key(id),
	foreign key (seller_id) references users(id)
)