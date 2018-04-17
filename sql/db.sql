CREATE TABLE text
(
   id character varying(100) PRIMARY KEY NOT NULL,
   editid character varying NOT NULL,
   passforview boolean NOT NULL,
   text character varying NOT NULl,
   created_at timestamp without time zone,
   title character varying,
   views integer
);