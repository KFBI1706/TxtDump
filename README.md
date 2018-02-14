# TxtDump


### Routes (Api):

#### Request a post:
```
HOST/api/v1/post/{id}/request | Returns the post id and content 
```
Example:
```json
curl http://localhost:1337/api/v1/post/{id}/request
```
Returns:
```json
{"PubID":9175728,"EditID":0,"Content":"I really REALLY hate javascript","Title":"Dette e ein title","Sucsess":true,"Time":"2017-12-27T00:00:00Z"}
```
#### Create Post:

```
HOST/api/v1/post/create | Creates the postk:
```
Example:
```json
curl -H "Content-Type: application/json" -X POST -d '{"Title":"Title","Content":"text"}' http://localhost:1337/api/v1/post/create
```
Response:
```json
{"PubID":5580586,"EditID":4553279,"Content":"text","Title":"Title","Sucsess":true,"Time":"0001-01-01T00:00:00Z"}
```
#### Edit Post:
```json
curl -H "Content-Type: application/json" -X POST -d '{"Title":"lmao",
"Content":"tyest"}' http://localhost:1337/api/v1/post/4750794/edit/8986640
```
#### Delete Post:
```
curl http://localhost:1337/api/v1/post/4750794/delete/8986640
```
<br>

#### Dbstring example:
The program looks for a file named dbstring when running this is then converted into the info used to connect to the DB for more info about this read: https://godoc.org/github.com/lib/pq
```
user=postgres dbname=db password=12345 host=HOSTIP
```