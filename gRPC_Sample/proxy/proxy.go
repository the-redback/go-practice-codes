package proxy

import (
	"flag"
	"net/http"
	"github.com/golang/glog"
	"log"

	"golang.org/x/net/context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "GoglandProjects/gRPC_Sample/exampleMessage"


)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:50053", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8088", mux)
}

func Call() {
	flag.Parse()
	defer glog.Flush()

	log.Println("Proxy Server running at port 8088")

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
