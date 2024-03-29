package main

import (
	"fmt"
        "context"
        pb "github.com/morix1500/now-grpc/proto"
        "google.golang.org/grpc"
        "net"
	"net/http"
)

type HelloService struct{}

func (h HelloService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
        return &pb.HelloResponse{
                Message: "Hello, " + in.Name,
        }, nil
}

func Hello(w http.ResponseWriter, r *http.Request) {
        s := grpc.NewServer()
        pb.RegisterHelloServiceServer(s, HelloService{})

        lis, err := net.Listen("tcp", ":5000")
        if err != nil {
                panic(err)
        }
	fmt.Println("start server...")
        if err := s.Serve(lis); err != nil {
                panic(err)
        }
	s.ServeHTTP(w, r)
	fmt.Println("start http server...")
}

//func main() {
//        s := grpc.NewServer()
//        pb.RegisterHelloServiceServer(s, HelloService{})
//
//        lis, err := net.Listen("tcp", ":5000")
//        if err != nil {
//                panic(err)
//        }
//	fmt.Println("start server...")
//        if err := s.Serve(lis); err != nil {
//                panic(err)
//        }
//}
