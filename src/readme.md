Dbstring example:
```
user=postgres dbname=db password=12345 host=HOSTIP
```

```SQL
CREATE TABLE text (
    id serial PRIMARY KEY,
    pubid integer NOT NULL,
    text varchar NOT NULL,
    created_at date
);
```