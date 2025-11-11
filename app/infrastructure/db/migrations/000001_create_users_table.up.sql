create table users
(
	id char(26) not null,
	email char(255) not null,
	name char(255) not null,
	hashed_password char(255) not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp on update current_timestamp,
	primary key(id)
);

create index idx_email on users (email);
