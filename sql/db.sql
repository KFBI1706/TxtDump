CREATE TABLE public.text
(
   id character varying(100) PRIMARY KEY NOT NULL,
   editid character varying NOT NULL,
   passforview boolean NOT NULL
   text character varying NOT NULL,
   created_at timestamp without time zone,
   title character varying,
   views integer,
);