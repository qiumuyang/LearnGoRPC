package main

import (
	"flag"
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"time"
)

func main() {
	// parse command line arguments
	var host string
	var token string
	flag.StringVar(&host, "h", "localhost:8088", "host:port")
	flag.StringVar(&token, "t", "123456", "auth token")
	flag.Parse()

	// connect to server
	client, err := jsonrpc.Dial("tcp", host)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	for {
		// make request
		request := TimeServiceRequest{AuthToken: token}
		response := new(TimeServiceResponse)
		err = client.Call("TimeService.GetTime", &request, &response)
		if err != nil {
			log.Fatal("TimeService error:", err)
		}

		// display response
		if response.Status != "ok" {
			fmt.Printf("error: %v\n", response.Status)
		} else {
			fmt.Printf("\r" + time.Unix(response.Time, 0).Format("15:04:05"))
		}
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	client.Close()
}
