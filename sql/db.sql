CREATE TABLE text (
    id serial PRIMARY KEY,
    pubid integer NOT NULL,
    text varchar NOT NULL,
    title varchar,
    created_at date
);