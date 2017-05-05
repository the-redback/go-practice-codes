server.go opens port at 50053 for gprc, 
and proxy.go opens 8088 proxy port by gPRC_gateway

proxy.go is runned by server.go as go routine 


1. command: 
cd $GOPATH/src/GoglandProjects/gRPC_Sample

2. command:  [this generates go file from proto file]

protoc -I /usr/local/include -I . \
       -I {$GOPATH}/src/github.com/googleapis/googleapis/ \
       --go_out=plugins=grpc:. \
       --grpc-gateway_out=logtostderr=true:. \
       exampleMessage/exampleMessage.proto


3.command:
cd $GOPATH/src/GoglandProjects/gRPC_Sample/exampleServer

4.command
go run server.go


5. keep this terminal open and open another terminal
6. command:

cd $GOPATH/go/src/GoglandProjects/gRPC_Sample/exampleClient

7.command:
go run client.go

8. give one string and two int input


or instead of 6,7,8 , run command:

curl --data "{\"name\": \"pc1\", \"a\":2,\"b\":3}" http://localhost:8088/echo




repeat 5,6,7,8 multiple times

curl --data "{\"name\": \"pc1\", \"a\":2,\"b\":3}" http://localhost:8088/echo