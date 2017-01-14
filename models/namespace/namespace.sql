CREATE TABLE IF NOT EXISTS public.namespace
(
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  label character varying,
  user_id uuid NOT NULL references users(user_id),
  created timestamp without time zone NOT NULL DEFAULT statement_timestamp(),
  active boolean NOT NULL DEFAULT false,
  removed boolean NOT NULL DEFAULT false,
  kube_exist boolean NOT NULL DEFAULT false,
  CONSTRAINT namespace_pkey PRIMARY KEY (id)
);
