package main

import (
	"net"
	"log"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"golang.org/x/net/context"

	pb "GoglandProjects/gRPC_Sample/exampleMessage"

	"GoglandProjects/gRPC_Sample/proxy"
)

const (
	port = ":50053"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	c:=in.A+in.B
	log.Println("Request from :"+in.Name+", where numbers are",in.A,in.B)

	return &pb.HelloReply{ "Successful request",c}, nil
}



func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Server running at port 50053")

	go proxy.Call() //gRPC gateway

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
