CREATE TABLE text
(
   id character varying(100) PRIMARY KEY NOT NULL,
   editid character varying NOT NULL,
   postperms integer NOT NULL,
   text character varying NOT NULl,
   created_at timestamp without time zone,
   title character varying,
   views integer
);
/*
# Note on post perms:
    1 = Anarchy anyone can view and edit without password
    2 = Normal/defualt anyone can view but password is needed for edits
    3 = Same as normal but password is required for viewing
*/