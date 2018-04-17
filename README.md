# TxtDump
## All current routes:
```
/
/api/v1/post/amount
/api/v1/post/{id}/request
/api/v1/post/create
/api/v1/post/{id}/edit/{editid}
/api/v1/post/{id}/delete/{editid}
/post/{id}/request
/post/{id}/edit/{editid}
/post/{id}/edit/{editid}/post
/post/{id}/delete/{editid}/
/post/create
/post/create/new
/documentation
/static/
```
## Routes (Api):
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
{"ID":9175728,"EditID":0,"Content":"I really REALLY hate javascript","Title":"Dette e ein title","Sucsess":true,"Time":"2017-12-27T00:00:00Z"}
```
#### Create Post:
```
HOST/api/v1/post/create | Creates the post:
```
Example:
```json
curl -H "Content-Type: application/json" -X POST -d '{"Title":"Title","Content":"text","EditID":"password"}' http://localhost:1337/api/v1/post/create
```
Response:
```json
{"ID":5580586,"Content":"text","Title":"Title","Sucsess":true,"Time":"0001-01-01T00:00:00Z"}
```
#### Edit Post:
```
HOST/api/v1/post/{ID} | Edits the post:
```
```json
curl -H "Content-Type: application/json" -X POST -d '{"Title":"lmao",
"Content":"tyest","EditID":"password"}' http://localhost:1337/api/v1/post/4750794/edit
```
#### Delete Post:
```
HOST/api/v1/post/{ID}/delete | Deletes the post:
```
Example:
```json
curl -H "Content-Type: application/json" -X POST -d '{"EditID":"password"}' http://localhost:1337/api/v1/post/4750794/delete
```
#### Dbstring example:
The program looks for a file named dbstring when running this is then converted into the info used to connect to the DB for more info about this read: https://godoc.org/github.com/lib/pq
```
user=postgres dbname=db password=12345 host=HOSTIP sslmode=disable
```
## Command line arguments
```
-setupdb Creates the table used for storing posts using info from dbstring file
-dropdb Drops the text table and all data. if run together with -setupdb it will drop then create a new emtpy table
-port Run with custom port. Defualt port is: 1337
```