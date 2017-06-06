package main

import (
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "GoglandProjects/gRPC_Sample/exampleMessage"
	"fmt"
)

const (
	address     = "localhost:50053"
	defaultName = "world"
)



func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	var name string
	fmt.Scanf("%s",&name)
	var a,b int32
	fmt.Scanf("%d",&a)
	fmt.Scanf("%d",&b)

	/*name :="asd"
	a:=4
	b:=6
*/
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{ name,int32(a),int32(b)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Message,r.Result)

}
