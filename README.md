# TxtDump
## All current routes:
```
/
/api/v1/post/amount
/api/v1/post/{id}/request
/api/v1/post/{id}/request
/api/v1/post/create
/api/v1/post/{id}/edit
/api/v1/post/{id}/delete
/post/{id}/request
/post/{id}/request/decrypt
/post/{id}/edit
/post/{id}/edit/decrypt
/post/{id}/edit/post
/post/{id}/delete
/post/{id}/delete/post
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
{"ID":9175728,"Hash":0,"Content":"I really REALLY hate javascript","Title":"Dette e ein title","Sucsess":true,"Time":"2017-12-27T00:00:00Z"}
```
#### Create Post:
```
HOST/api/v1/post/create | Creates the post:
```
Example:
```json
curl -H "Content-Type: application/json" -X POST -d '{"Title":"Title","Content":"text","Hash":"password"}' http://localhost:1337/api/v1/post/create
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
curl -H "Content-Type: application/json" -X POST -d '{"Title":"lmao", "Content":"tyest","Hash":"password"}' http://localhost:1337/api/v1/post/4750794/edit
```
#### Delete Post:
```
HOST/api/v1/post/{ID}/delete | Deletes the post:
```
Example:
```json
curl -H "Content-Type: application/json" -X POST -d '{"Hash":"password"}' http://localhost:1337/api/v1/post/4750794/delete
```
#### Dbstring example: 
The program looks for a file named dbstring when running this is then converted into the info used to connect to the DB for more info about this read: https://godoc.org/github.com/lib/pq
```
user=postgres dbname=web password=12345 host=192.168.10.179
```
## Command line arguments
```
-setupdb Creates the table used for storing posts using info from dbstring file
-dropdb Drops the text table and all data. if run together with -setupdb it will drop then create a new emtpy table
-port Run with custom port. Defualt port is: 1337
```
## config.json
Change "Path" to your own path for the project
```json
"Path": "/home/vetlo/Documents/code/go/src/github.com/KFBI1706/TxtDump/"
```