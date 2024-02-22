package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_test/pb"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

var (
	GrpcMap map[string]*grpc.ClientConn
	mutex   sync.Mutex
)

func InitGrpc() {
	GrpcMap = make(map[string]*grpc.ClientConn, 2)
}

func GetGrpcConn(url string) (*grpc.ClientConn, error) {
	conn := GrpcMap[url]
	if conn == nil {
		var err error
		mutex.Lock()
		defer mutex.Unlock()
		log.Println("连接服务器")
		c, err := credentials.NewClientTLSFromFile(`D:\code_info\go_project\GO\gotest\grpc\secret\5361280__yebaojiasu.com.pem`, "yebaojiasu.com")
		//conn, err = grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		conn, err = grpc.Dial(*addr, grpc.WithTransportCredentials(c))
		if err != nil {
			return nil, err
		}
		GrpcMap[url] = conn
	}
	return conn, nil
}

func main() {
	flag.Parse()
	InitGrpc()
	// Set up a connection to the server.
	conn, err := GetGrpcConn("url1")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			name := fmt.Sprintf("name%d", r.Intn(10000))
			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
			if err != nil {
				log.Printf("请求失败: %v", err)
			}
			log.Printf("Greeting: %s", r.GetMessage())
			cancel()
		}
	}
}

var client *grpc.ClientConn

func InitClient() {
	if client == nil {
		//client
	}
}
