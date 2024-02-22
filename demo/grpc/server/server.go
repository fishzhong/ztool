// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_test/pb"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("ReceivedName: %v", in.GetName())

	//return nil, errors.New("ces")

	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var c credentials.TransportCredentials
	c, err = credentials.NewServerTLSFromFile(
		`D:\code_info\go_project\GO\gotest\grpc\secret\5361280__yebaojiasu.com.pem`,
		`D:\code_info\go_project\GO\gotest\grpc\secret\5361280__yebaojiasu.com.key`,
	)
	//altsTC := alts.NewServerCreds(alts.DefaultServerOptions())

	//c2, err := oauth.NewJWTAccessFromKey([]byte(""))
	s := grpc.NewServer(grpc.Creds(c))
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}

	//fmt.Println("Shutting down gRPC server...")
	//s.GracefulStop()
	//fmt.Println("gRPC server stopped.")
}
