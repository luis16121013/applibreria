create database soft;

--TABLE ROLE USERS
create table roleUserTable(
	id_role int  primary key auto_increment,
	descripcion varchar(30)
);
--TABLE USERS
create table usersTable(
	id_user int primary key auto_increment,
	nombre varchar(50),
	apellido varchar(50),
	dni varchar(8),
	tipo_user int,
	foreign key (tipo_user) references roleUserTable(id_role)
);

--TABLE ACCESS USERS 
create table loginTable(
	id_user int,
	usuario varchar(20) not null,
	pwd varchar(100) not null,
	foreign key (id_user) references usersTable(id_user)
);

--TABLE BOOK
create table bookTable(
	id_libro int primary key auto_increment,
	titulo varchar(100),
	autor varchar(100),
	editorial varchar(100),
	edicion varchar(100),
	pais varchar(100),
	ciudad varchar(100),
	fechaPublicacion datetime
);

--TABLE CATEGORY
create table categoryTable(
	id_categoria int primary key auto_increment,
	descripcion varchar(100)
);
--TABLE PRODUCT
create table productTable(
	id_product int primary key auto_increment,
	nombre varchar(100),
	descripcion varchar(100),
	id_categoria int,
	precio float,
	stock int,
	foreign key (id_categoria) references categoryTable(id_categoria)
);

--TABLE SALE
create table saleTable(
	id_venta int primary key auto_increment,
	id_user int,
	id_product int,
	fecha datetime,
	foreign key(id_user) references usersTable(id_user),
	foreign key(id_product) references productTable(id_product)
);
