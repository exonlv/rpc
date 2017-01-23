create extension "uuid-ossp";
CREATE TABLE IF NOT EXISTS tcp(
	id 		uuid NOT NULL DEFAULT uuid_generate_v4(),
	user_id uuid NOT NULL,
	channel varchar NOT NULL,
	active	timestamp,
	opened	boolean,
	ip		varchar(15) NOT NULL
);
