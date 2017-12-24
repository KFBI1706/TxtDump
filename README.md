# TxtDump


### Routes:

#### Request a post:
```
HOST/post/{post id}/request | Returns the post id and content 
```
Example:
```
curl http://localhost:1337/post/11/request
{"ID":11,"Content":"not implemented yet"}
```
#### Request post ID: 
not completly sure this is needed but its implemented for now

```
HOST/random/test
```
Example:
```
curl http://localhost:1337/random/test         
{"ID":3652446,"Content":"Your ID"}
```

#### Create Post with requested ID:

```
HOST/post/create | Creates the post with the submitted ID:
```
Example:
```
curl -H "Content-Type: application/json" -X POST -d '{"ID":23,"Content":"Doope"}' http://localhost:1337/post/create 
Returns: {"ID":23,"Content":"Doope","Sucsess":true}   
```

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