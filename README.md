# TxtDump


### Routes:

#### Request a post:
```
HOST/post/{post id}/request | Returns the post id and content 
```
Example:
```json
curl http://localhost:1337/post/11/request
{"ID":11,"Content":"not implemented yet"}
```
#### Request post ID: 
not completly sure this is needed but its implemented for now

```
HOST/random/test
```
Example:
```json
curl http://localhost:1337/random/test         
{"ID":3652446,"Content":"Your ID"}
```

#### Create Post with requested ID:

```
HOST/post/create | Creates the post with the submitted ID:
```
Example:
```json
curl -H "Content-Type: application/json" -X POST -d '{"Content":"I really hate javascript"}' http://localhost:1337/post/create 
Returns: {"PubID":9175728,"Content":"I really hate Javascript","Sucsess":true,"Time":""}
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