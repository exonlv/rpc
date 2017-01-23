CREATE TABLE IF NOT EXISTS token(
	token		uuid NOT NULL DEFAULT uuid_generate_v4(),
	user_id uuid REFERENCES public.users (user_id),
	created timestamp without time zone NOT NULL DEFAULT now(),
	expired	timestamp without time zone NOT NULL DEFAULT now() + interval '1 day' * 90,
	active boolean NOT NULL DEFAULT true
);
