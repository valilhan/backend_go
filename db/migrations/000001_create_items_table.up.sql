CREATE TABLE IF NOT EXISTS main_balance (
	user_id  varchar not null  PRIMARY KEY,
	balance FLOAT CHECK (balance >= 0),
	created_at TIMESTAMPtz not null DEFAULT(now())
);

CREATE TABLE IF NOT EXISTS reserve_balance (
	id serial not null  PRIMARY key,
	user_id varchar not null,
	service_id varchar not null,
	order_id varchar not null UNIQUE,
	price float not null,
	created_at TIMESTAMPtz not null DEFAULT(now())
);

CREATE TABLE IF NOT EXISTS revenue (
	id serial not null PRIMARY key,
	user_id varchar not null,
	service_id varchar not null,
	order_id varchar not null,
	price float not null,
	created_at TIMESTAMPtz not null DEFAULT(now())
);

Alter table reserve_balance add foreign key (user_id) references main_balance(user_id);
alter table revenue add foreign key(user_id) references main_balance(user_id);
create index on main_balance (user_id);
create index on reserve_balance(user_id);
create index on reserve_balance(service_id);
create index on reserve_balance(order_id);
create index on revenue(user_id);
create index on revenue(service_id);
create index on revenue(order_id);