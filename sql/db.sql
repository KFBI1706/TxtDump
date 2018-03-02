CREATE TABLE public.text
(
   id character varying(100) PRIMARY KEY NOT NULL,
   editid integer,
   text character varying NOT NULL,
   created_at timestamp without time zone,
   title character varying
);