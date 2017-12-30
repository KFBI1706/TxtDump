# TxtDump


### Routes (Api):

#### Request a post:
```
HOST/api/v1/post/{id}/request | Returns the post id and content 
```
Example:
```json
curl http://localhost:1337/api/v1/post/{id}/request
{"ID":11,"Content":"not implemented yet"}
```
#### Create Post with requested ID:

```
HOST/api/v1/post/create | Creates the post with the submitted ID:
```
Example:
```json
curl -H "Content-Type: application/json" -X POST -d '{"Content":"I really hate javascript"}' http://localhost:1337/api/v1/post/create 
Returns: {"PubID":9175728,"Content":"I really hate Javascript","Sucsess":true,"Time":""}
```

Dbstring example:
The program looks for a file named dbstring when running this is then converted into the info used to connect to the DB for more info about this read: https://godoc.org/github.com/lib/pq
```
user=postgres dbname=db password=12345 host=HOSTIP
```

```SQL
CREATE TABLE text (
    id serial PRIMARY KEY,
    pubid integer NOT NULL,
    text varchar NOT NULL,
    title varchar,
    created_at date
);
```