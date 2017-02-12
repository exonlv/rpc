
CREATE TABLE IF NOT EXISTS public.volumes
(
  volume_id uuid PRIMARY KEY unique NOT NULL DEFAULT uuid_generate_v4(),
  label character varying UNIQUE,
  replicas integer NOT NULL CHECK (replicas >= 2 AND replicas <= 5),
  volumeservers uuid[],
  limits integer,
  user_id uuid UNIQUE NOT NULL references users(user_id),
  created timestamp without time zone NOT NULL DEFAULT statement_timestamp(),
  active boolean NOT NULL DEFAULT true,
  exists boolean NOT NULL DEFAULT false
);