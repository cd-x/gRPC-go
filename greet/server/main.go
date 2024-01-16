package main
import (
	"log"
	"net"
	"google.golang.org/grpc"
	pb "github.com/cd-x/grpc-go/greet/proto"
)
type Server struct{
	pb.GreetServiceServer
}
var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Failed to listen on: %v\n",err)
	}

	log.Printf("Listening on: %v\n",addr)
	s:= grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v\n", err);
	}
}