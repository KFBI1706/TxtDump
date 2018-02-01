CREATE TABLE public.text
(
  id integer NOT NULL DEFAULT nextval('text_id_seq'::regclass),
  pubid integer NOT NULL,
  text character varying NOT NULL,
  created_at date,
  title character varying,
  CONSTRAINT text_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);