create schema if not exists dev_book;

create sequence dev_book.user_seq;

CREATE TABLE "dev_book".user (
	id int4 NOT NULL DEFAULT nextval('"dev_book".user_seq'::regclass),
	name varchar(100) not null,
	nick varchar(50) not null,
	email varchar(100) not null, 
	password varchar(45) not null, 
	cretae_at timestamp default current_timestamp,
	CONSTRAINT id_user_pk PRIMARY KEY (id)
);
