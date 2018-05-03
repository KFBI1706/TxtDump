#/api/v1/post/amount
#/api/v1/post/{id}/request
#/api/v1/post/create
#/api/v1/post/{id}/edit/{editid}
#/api/v1/post/{id}/delete/{editid}
HOST="localhost:1337"
ID="4060228"
#curl -H "Content-Type: application/json" -X POST -d '{"Title":"Title","Content":"text","Hash":"password"}' $HOST/api/v1/post/create
curl $HOST/api/v1/post/$ID/request
#curl -H "Content-Type: application/json" -X POST -d '{"Title":"lmao", "Content":"tyest","Hash":"password"}' $HOST/api/v1/post/4750794/edit
#curl -H "Content-Type: application/json" -X POST -d '{"Hash":"password"}' $HOST/api/v1/post/4750794/delete
