/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"log"
	"net"

	kek "service_db_handler/protobuf/kek/gen/kek"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	kek.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *kek.HelloRequest) (*kek.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &kek.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "192.168.50.20:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	kek.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// package main

// import (
// 	"fmt"
// 	"net"
// )

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	// Read data from the client
// 	buffer := make([]byte, 1024)
// 	_, err := conn.Read(buffer)
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 		return
// 	}

// 	// Print received message
// 	fmt.Println("Received message:", string(buffer))

// 	// Respond to the client
// 	conn.Write([]byte("Message received!"))
// }

// func main() {
// 	listener, err := net.Listen("tcp", "192.168.50.20:8080")
// 	if err != nil {
// 		fmt.Println("Error listening:", err.Error())
// 		return
// 	}
// 	defer listener.Close()

// 	fmt.Println("Server started. Waiting for connections...")

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("Error accepting:", err.Error())
// 			return
// 		}
// 		fmt.Println("Client connected:", conn.RemoteAddr())

// 		// Handle client's connection concurrently
// 		go handleConnection(conn)
// 	}
// }
