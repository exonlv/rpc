create extension "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.users
(
  user_id uuid NOT NULL DEFAULT uuid_generate_v4(),
  login character varying NOT NULL,
  pwd_key varchar NOT NULL, --len(base64(seq [32]byte)) == 128
  salt varchar NOT NULL,
  name character varying NOT NULL,
  last_name character varying,
  email character varying NOT NULL,
  active boolean NOT NULL DEFAULT false,
  notes character varying,
  register_date timestamp without time zone NOT NULL DEFAULT statement_timestamp(),
  CONSTRAINT user_id PRIMARY KEY (user_id)
);
