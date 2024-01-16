package main
import (
	"log"
	"net"
	pb "github.com/cd-x/grpc-go/greet/proto"
)
var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Failed to listen on: %v\n",err)
	}

	log.Printf("Listening on: %v\n",addr)
}