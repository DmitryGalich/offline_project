// package main

// import (
// 	"context"
// 	"flag"
// 	"log"
// 	kek "service_db_handler/protobuf/kek/gen/kek"
// 	"time"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// const (
// 	defaultName = "world"
// )

// var (
// 	addr = flag.String("addr", "localhost:50051", "the address to connect to")
// 	name = flag.String("name", defaultName, "Name to greet")
// )

// func main() {
// 	flag.Parse()

// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()

// 	c := kek.NewGreeterClient(conn)

// 	// Contact the server and print out its response.
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	r, err := c.SayHello(ctx, &kek.HelloRequest{Name: *name})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("Greeting: %s", r.GetMessage())
// }

package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	// Send data to the server
	message := "Hello, server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}

	// Read response from the server
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}

	// Print server's response
	fmt.Println("Server says:", string(buffer))
}
