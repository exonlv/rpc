CREATE TABLE IF NOT EXISTS public.volumeservers
(
  volumeserver_id uuid NOT NULL DEFAULT uuid_generate_v4(),
  ip character varying,
  path character varying,
  memory integer,
  created timestamp without time zone NOT NULL DEFAULT statement_timestamp(),
  active boolean NOT NULL DEFAULT true,
  groups varchar NOT NULL DEFAULT 'default',
  disk_type varchar NOT NULL DEFAULT 'HDD'
);
