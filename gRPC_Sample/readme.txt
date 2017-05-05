1. command: 
cd /home/maruf/go/src/GoglandProjects/gRPC_Sample

2. command:  [this generates go file from proto file]
protoc -I exampleMessage exampleMessage/exampleMessage.proto --go_out=plugins=grpc:exampleMessage



3.command:
cd /home/maruf/go/src/GoglandProjects/gRPC_Sample/exampleServer

4.command
go run server.go


5. keep this terminal open and open another terminal
6. command:

cd /home/maruf/go/src/GoglandProjects/gRPC_Sample/exampleClient

7.command:
go run client.go

8. give one string and two int input

repeat 5,6,7,8 multiple times