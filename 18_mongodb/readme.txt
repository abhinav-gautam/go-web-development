Get drivers for mongo
go get gopkg.in/mgo.v2
go get gopkg.in/mgo.v2/bson

Curl commands
curl http://localhost:8080/user/<enter-user-id-here>
curl -X POST -H "Content-Type: application/json" -d '{"name":"Abhinav","gender":"male","age":27}' http://localhost:8080/user
curl -X DELETE http://localhost:8080/user/<enter-user-id-here>