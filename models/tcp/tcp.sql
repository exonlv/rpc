CREATE TABLE IF NOT EXISTS tcp(
	id 		serial primary key,
	user_id int NOT NULL,
	channel varchar NOT NULL,
	active	timestamp,
	opened	boolean,
	ip		varchar(15) NOT NULL
);
