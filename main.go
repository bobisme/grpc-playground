package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	// "github.com/golang/glog"
	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	// echosvc "path/to/your_service_package"
	echosvc "github.com/bobisme/grpc-playground/gen/pb-go"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

type echoServer struct{}

func (e *echoServer) Greet(
	c context.Context, req *echosvc.GreetingRequest) (
	*echosvc.GreetingResponse, error) {
	//
	return &echosvc.GreetingResponse{
		Msg: fmt.Sprintf("%s, %s", req.Type, req.Name),
	}, nil
}

func (e *echoServer) Echo(c context.Context, t *echosvc.EchoMessage) (*echosvc.EchoMessage, error) {
	fmt.Println("got", t)
	return t, nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithInsecure()}
	// err := echosvc.RegisterEchoServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	// if err != nil {
	// 	return err
	// }
	// return http.ListenAndServe(":8080", mux)

	lis, err := net.Listen("tcp", *echoEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	echosvc.RegisterEchoServiceServer(grpcServer, &echoServer{})
	// determine whether to use TLS
	return grpcServer.Serve(lis)
}

func main() {
	flag.Parse()
	// defer log.Flush()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
