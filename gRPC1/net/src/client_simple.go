package src

import (
	"fmt"
	"log"
	"net"
	pb "./src/simple"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)
type Listener int
func (l *Listener) GetLine(ctx context.Context, in *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	rv := in.Data
	fmt.Printf("Receive: %v\n", rv)
	return &pb.SimpleResponse{Data: rv}, nil
}
func main() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:12345")
	if err != nil {
		log.Fatal(err)
	}
	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	listener := new(Listener)
	pb.RegisterSimpleServer(s, listener)
	s.Serve(inbound)
}

import (
"bufio"
"log"
"os"
pb "./src/simple"
"golang.org/x/net/context"
"google.golang.org/grpc"
)
func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := pb.NewSimpleClient(conn)
	in := bufio.NewReader(os.Stdin)
	for {
		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		reply, err := c.GetLine(context.Background(), &pb.SimpleRequest{Data: string(line)})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reply: %v, Data: %v", reply, reply.Data)
	}
}