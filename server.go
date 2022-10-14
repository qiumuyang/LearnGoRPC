package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// parse command line arguments
	var port string
	var tokenFile string
	flag.StringVar(&port, "p", "8088", "port")
	flag.StringVar(&tokenFile, "t", "tokens.txt", "token file")
	flag.Parse()

	// create authentication
	auth := &FileBasedAuthentication{Tokens: []string{"123456"}}
	auth.LoadTokens(tokenFile)

	// register time service
	rpc.Register(&TimeService{Auth: auth})
	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	} else {
		log.Println("Listening on port", port)
	}

	// accept connections
	defer socket.Close()
	for {
		conn, err := socket.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()
		go jsonrpc.ServeConn(conn)
	}
}
